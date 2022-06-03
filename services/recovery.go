package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"rarity-backend/constants"
	"rarity-backend/dlt"
	"rarity-backend/handlers"
	"rarity-backend/helpers"
	"rarity-backend/metadata"
	"rarity-backend/models"
	"rarity-backend/store"
	"rarity-backend/structs"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/bson"
)

// RecoverProcess is the main function which handles the polling and processing of mint and morph events
func RecoverProcess(ethClient *dlt.EthereumClient, contractAbi abi.ABI, instance *store.Store, address string, configService *structs.ConfigService,
	dbInfo structs.DBInfo, txState map[string]map[uint]bool, morphCostMap map[string]float32) {
	var wg sync.WaitGroup
	mintsMutex := structs.MintsMutex{TokensMap: make(map[string]bool)}
	eventLogsMutex := structs.EventLogsMutex{EventLogs: []types.Log{}}
	genesMap := make(map[string]string)
	tokenToMorphEvent := make(map[string]types.Log)

	lastProcessedBlockNumber := collectEvents(ethClient, contractAbi, instance, address, configService, dbInfo.PolymorphDBName, dbInfo.RarityCollectionName, dbInfo.BlocksCollectionName, 0, 0, &wg, &eventLogsMutex)

	mintsLogs := make([]types.Log, 0)

	// Persist mints
	for _, ethLog := range eventLogsMutex.EventLogs {
		eventSig := ethLog.Topics[0].String()
		switch eventSig {
		case constants.MintEvent.Signature:
			//if !txState[ethLog.TxHash.Hex()][ethLog.Index] {
			//	txMap := make(map[uint]bool)
			//	txMap[ethLog.Index] = true
			//	txState[ethLog.TxHash.Hex()] = txMap
			//
			//
			//}
			wg.Add(1)
			go processMint(ethLog, &wg, contractAbi, configService, dbInfo.PolymorphDBName, dbInfo.RarityCollectionName, &mintsMutex)
			mintsLogs = append(mintsLogs, ethLog)
		}
	}

	wg.Wait()
	if len(mintsMutex.Documents) > 0 {
		fmt.Println(mintsMutex.Documents)
		handlers.PersistMintEvents(mintsMutex.Documents, dbInfo.PolymorphDBName, dbInfo.RarityCollectionName)
		handlers.DeleteV1Rarity(dbInfo.PolymorphDBName, &mintsLogs)
	}

	if len(mintsMutex.Transactions) > 0 {
		handlers.SaveTransactions(mintsMutex.Transactions, dbInfo.PolymorphDBName, dbInfo.TransactionsCollectionName)
	}

	// Sort polymorphs
	helpers.SortMorphEvents(eventLogsMutex.EventLogs)
	// Persist Morphs
	for _, ethLog := range eventLogsMutex.EventLogs {
		eventSig := ethLog.Topics[0].String()
		switch eventSig {
		case constants.MorphEvent.Signature:
			processInitialMorphs(ethLog, ethClient, contractAbi, instance, configService, dbInfo, txState, genesMap, tokenToMorphEvent, morphCostMap)
		}
	}

	// Persist final scrambles
	for id := range genesMap {
		ethLog := tokenToMorphEvent[id]
		processFinalMorphs(ethLog, ethClient, contractAbi, instance, configService, dbInfo, txState, genesMap, morphCostMap)
	}

	// Persist Ranking
	handlers.UpdateAllRanking(dbInfo.PolymorphDBName, dbInfo.RarityCollectionName)
	// Persist block
	res, err := handlers.CreateOrUpdateLastProcessedBlock(lastProcessedBlockNumber, dbInfo.PolymorphDBName, dbInfo.BlocksCollectionName)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(res)
	}
}

// processMint is the core function for processing mint events metadata. It unpacks event data, calculates rarity score, prepares database entity but doesn't persist it
//
// Uses Mutes and WaitGroup in order to process events faster and prevent race conditions.
func processMint(mintEvent types.Log, wg *sync.WaitGroup, contractAbi abi.ABI, configService *structs.ConfigService, polymorphDBName string, rarityCollectionName string, mintsMutex *structs.MintsMutex) {
	defer wg.Done()
	var event structs.PolymorphEvent
	mintsMutex.Mutex.Lock()
	contractAbi.UnpackIntoInterface(&event, constants.MintEvent.Name, mintEvent.Data)
	event.MorphId = mintEvent.Topics[1].Big()
	event.OldGene = big.NewInt(0)
	if event.NewGene.String() != "0" && !mintsMutex.TokensMap[event.MorphId.String()] {
		g := metadata.Genome(event.NewGene.String())
		metadataJson := (&g).Metadata(event.MorphId.String(), configService)
		rarityResult := CalulateRarityScore(metadataJson.Attributes, true)
		mintEntity := helpers.CreateMorphEntity(event, metadataJson, true, rarityResult)

		mintsMutex.Mints = append(mintsMutex.Mints, mintEntity)
		mintsMutex.TokensMap[event.MorphId.String()] = true
		var bdoc interface{}
		jsonEntity, _ := json.Marshal(mintEntity)
		bson.UnmarshalExtJSON(jsonEntity, false, &bdoc)
		mintsMutex.Documents = append(mintsMutex.Documents, bdoc)

		transaction := models.Transaction{
			BlockNumber: mintEvent.BlockNumber,
			TxIndex:     mintEvent.TxIndex,
			TxHash:      mintEvent.TxHash.Hex(),
			LogIndex:    mintEvent.Index,
		}
		var txBdoc interface{}
		jsonTx, _ := json.Marshal(transaction)
		bson.UnmarshalExtJSON(jsonTx, false, &txBdoc)

		mintsMutex.Transactions = append(mintsMutex.Transactions, txBdoc)

		// go handlers.SaveTransaction(dbInfo.PolymorphDBName, dbInfo.TransactionsCollectionName, models.Transaction{
		// 	BlockNumber: ethLog.BlockNumber,
		// 	TxIndex:     ethLog.TxIndex,
		// 	TxHash:      ethLog.TxHash.Hex(),
		// 	LogIndex:    ethLog.Index,
		// })

	} else {
		log.Println("Empty gene mint event for morph id: " + event.MorphId.String())
	}
	mintsMutex.Mutex.Unlock()
}

// processInitialMorphs is the core function for processing morph events. It's contains the trickiest logic in the app because TokenMorphed event emits the old gene in both the new gene and old gene parameters.
//
// We're interested in morph events with event type 1. (0 is Morph, 2 is Transfer)
//
// We can't be sure how many morph events each polymorph has. This is why we have to process the morph event only after we've got a chronological pair of genes.
// In oldGenesMap we keep track of the tokenId -> gene mappings.
//
// If there isn't already existing mapping for the token - we save the current gene in the mapping and proceed to saving the information for the newest gene received from the contract.
//
// If there is existing mapping - this means we've got a chronological pair of morph events of a polymorph and we can process them to find out which traits have changed.
// We compare the old and the new gene and create a history snapshot of the changes, persists the increment scramble/morph in the rarity collection and persists the event transaction in the transactions collection
//
// We save the new gene to the oldGenesMap and repeat the process for the next event for this polymorph.
//
// !! It's important to note which gene is passed as the new one and which as the old one in order to understand how the logic works.
func processInitialMorphs(morphEvent types.Log, ethClient *dlt.EthereumClient, contractAbi abi.ABI, instance *store.Store, configService *structs.ConfigService, dbInfo structs.DBInfo,
	txState map[string]map[uint]bool, oldGenesMap map[string]string, tokenToMorphEvent map[string]types.Log, morphCostMap map[string]float32) {
	var mEvent structs.MorphedEvent
	err := contractAbi.UnpackIntoInterface(&mEvent, constants.MorphEvent.Name, morphEvent.Data)
	if err != nil {
		log.Fatalln(err)
	}

	// 1 is Morph event
	txMap, hasTxMap := txState[morphEvent.TxHash.Hex()]
	if mEvent.EventType == 1 && (!hasTxMap || !txMap[morphEvent.Index]) {
		log.Println()
		log.Printf("\nBlock Num: %v\nTxIndex: %v\nEventIndex:%v\n", morphEvent.BlockNumber, morphEvent.TxIndex, morphEvent.Index)

		mId := morphEvent.Topics[1].Big()

		// This will get the newest gene
		result, err := instance.GeneOf(&bind.CallOpts{}, mId)
		if err != nil {
			log.Println(err)
		}
		mEvent.NewGene = result
		var geneDifferences, geneIdx int
		newAttr, oldAttr := structs.Attribute{}, structs.Attribute{}
		if oldGenesMap[mId.String()] != "" {
			geneIdx, geneDifferences = helpers.DetectGeneDifferences(oldGenesMap[mId.String()], mEvent.OldGene.String())
			if geneDifferences <= 2 {
				newAttr, oldAttr = helpers.GetAttribute(mEvent.OldGene.String(), oldGenesMap[mId.String()], geneIdx, configService)
			}
			block, err := ethClient.Client.BlockByNumber(context.Background(), big.NewInt(int64(morphEvent.BlockNumber)))
			if err != nil {
				log.Println(err)
			}
			polySnapshot := helpers.CreateMorphSnapshot(geneDifferences, mId.String(), mEvent.OldGene.String(), oldGenesMap[mId.String()], block.Time(), oldAttr, newAttr, morphCostMap, configService)
			go handlers.SavePolymorphHistory(polySnapshot, dbInfo.PolymorphDBName, dbInfo.HistoryCollectionName)
			go handlers.SaveMorphPrice(models.MorphCost{TokenId: mId.String(), Price: morphCostMap[mId.String()]}, dbInfo.PolymorphDBName, dbInfo.MorphCostCollectionName)
		}
		toSaveGene := oldGenesMap[mId.String()]
		oldGenesMap[mId.String()] = mEvent.OldGene.String()
		tokenToMorphEvent[mId.String()] = morphEvent

		g := metadata.Genome(mEvent.NewGene.String())
		metadataJson := (&g).Metadata(mId.String(), configService)

		rarityResult := CalulateRarityScore(metadataJson.Attributes, false)
		morphEntity := helpers.CreateMorphEntity(structs.PolymorphEvent{NewGene: mEvent.NewGene, OldGene: mEvent.OldGene, MorphId: mId}, metadataJson, false, rarityResult)

		res, err := handlers.PersistSinglePolymorph(morphEntity, dbInfo.PolymorphDBName, dbInfo.RarityCollectionName, toSaveGene, geneDifferences)
		if err != nil {
			log.Println(err)
		} else {
			log.Println(res)
		}

		if !hasTxMap {
			txMap = make(map[uint]bool)
			txState[morphEvent.TxHash.Hex()] = txMap
		}
		txState[morphEvent.TxHash.Hex()][morphEvent.Index] = true
		go handlers.SaveTransaction(dbInfo.PolymorphDBName, dbInfo.TransactionsCollectionName, models.Transaction{
			BlockNumber: morphEvent.BlockNumber,
			TxIndex:     morphEvent.TxIndex,
			TxHash:      morphEvent.TxHash.Hex(),
			LogIndex:    morphEvent.Index,
		})
	} else if txMap[morphEvent.Index] {
		log.Println("Already processed morph event! Skipping...")
	}
}

// processFinalMorphs is has almost the same logic as processInitialMorphs. It's idea is to process all the final mappings in oldGenesMap parameter.
//
// We're interested in morph events with event type 1 (0 is Morph, 2 is Transfer)
//
// What does this mean: At this point we've processed some morph events in processInitialMorphs but we still got some left in the oldGenesMap.
// Every gene in the map means that this is the latest morph event and must be compared with the current gene of the polymorph
//
// We compare the old and the new gene and create a history snapshot of the changes, persists the increment scramble/morph in the rarity collection.
//
// We don't persist the transaction as the transaction has already been persisted in processInitialMorphs.
//
// !! It's important to note which gene is passed as the new one and which as the old one in order to understand how the logic works.
func processFinalMorphs(morphEvent types.Log, ethClient *dlt.EthereumClient, contractAbi abi.ABI, instance *store.Store, configService *structs.ConfigService, dbInfo structs.DBInfo,
	txState map[string]map[uint]bool, oldGenesMap map[string]string, morphCostMap map[string]float32) {
	var mEvent structs.MorphedEvent
	err := contractAbi.UnpackIntoInterface(&mEvent, constants.MorphEvent.Name, morphEvent.Data)
	if err != nil {
		log.Fatalln(err)
	}

	mId := morphEvent.Topics[1].Big()

	// This will get the newest gene
	result, err := instance.GeneOf(&bind.CallOpts{}, mId)
	if err != nil {
		log.Println(err)
	}
	mEvent.NewGene = result

	g := metadata.Genome(mEvent.NewGene.String())
	genes := g.Genes()

	revGenes := metadata.ReverseGenesOrder(genes)

	b := strings.Builder{}
	b.WriteString(constants.POLYMORPH_IMAGE_URL)

	for _, gene := range revGenes {
		b.WriteString(gene)
	}
	b.WriteString(".jpg")

	// Currently, the front-end fetches imageURLs from the rarity-backend instead of from the Metadata API
	// So if the image doesn't exist, we query the metadata API to get it generated
	if !metadata.ImageExists(b.String()) {
		_, err := http.Get(constants.IMAGES_METADATA_URL + mId.String())
		if err != nil {
			log.Fatalf("Couldn't query images function. Original error: %v", err)
		} else {
			fmt.Println("Queried Metadata with link: ", constants.IMAGES_METADATA_URL+mId.String())
		}
	}

	newAttr, oldAttr := structs.Attribute{}, structs.Attribute{}
	geneIdx, geneDifferences := helpers.DetectGeneDifferences(oldGenesMap[mId.String()], mEvent.NewGene.String())
	if geneDifferences <= 2 {
		newAttr, oldAttr = helpers.GetAttribute(mEvent.NewGene.String(), oldGenesMap[mId.String()], geneIdx, configService)
	}
	block, err := ethClient.Client.HeaderByNumber(context.Background(), big.NewInt(int64(morphEvent.BlockNumber)))
	if err != nil {
		log.Println(err)
	}
	polySnapshot := helpers.CreateMorphSnapshot(geneDifferences, mId.String(), mEvent.NewGene.String(), oldGenesMap[mId.String()], block.Time, oldAttr, newAttr, morphCostMap, configService)
	go handlers.SavePolymorphHistory(polySnapshot, dbInfo.PolymorphDBName, dbInfo.HistoryCollectionName)
	go handlers.SaveMorphPrice(models.MorphCost{TokenId: mId.String(), Price: morphCostMap[mId.String()]}, dbInfo.PolymorphDBName, dbInfo.MorphCostCollectionName)

	g = metadata.Genome(mEvent.NewGene.String())
	metadata := (&g).Metadata(mId.String(), configService)

	rarityResult := CalulateRarityScore(metadata.Attributes, false)
	morphEntity := helpers.CreateMorphEntity(structs.PolymorphEvent{NewGene: mEvent.NewGene, MorphId: mId}, metadata, false, rarityResult)

	res, err := handlers.PersistSinglePolymorph(morphEntity, dbInfo.PolymorphDBName, dbInfo.RarityCollectionName, oldGenesMap[mId.String()], geneDifferences)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(res)
	}
}

// NOT BEING USED CURRENTLY
// This should be only used when you are GUARANTEED that there will be no more than one event for a single polymorph in each poll -> It writes incorrect data when there are > 1
// func processMorphs(morphEvent types.Log, wg *sync.WaitGroup, contractAbi abi.ABI, instance *store.Store, configService *structs.ConfigService, polymorphDBName string,
// 	rarityCollectionName string, transactionsCollectionName string, txState map[string]map[uint]bool, oldGenesMap map[string]string, tokenToMorphEvent map[string]types.Log) {
// 	var mEvent structs.MorphedEvent
// 	err := contractAbi.UnpackIntoInterface(&mEvent, constants.MorphEvent.Name, morphEvent.Data)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	// 1 is Morph event
// 	txMap, hasTxMap := txState[morphEvent.TxHash.Hex()]
// 	if mEvent.EventType == 1 && (!hasTxMap || !txMap[morphEvent.Index]) {
// 		log.Println()
// 		log.Printf("\nBlock Num: %v\nTxIndex: %v\nEventIndex:%v\n", morphEvent.BlockNumber, morphEvent.TxIndex, morphEvent.Index)

// 		mId := morphEvent.Topics[1].Big()
// 		// This will get the newest gene
// 		result, err := instance.GeneOf(&bind.CallOpts{}, mId)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		mEvent.NewGene = result
// 		geneDifferences := helpers.DetectGeneDifferences(mEvent.OldGene.String(), mEvent.NewGene.String())

// 		g := metadata.Genome(mEvent.NewGene.String())
// 		metadataJson := (&g).Metadata(mId.String(), configService)

// 		rarityResult := rarityIndex.CalulateRarityScore(metadataJson.Attributes, false)
// 		morphEntity := helpers.CreateMorphEntity(structs.PolymorphEvent{
// 			NewGene: mEvent.NewGene,
// 			OldGene: mEvent.OldGene,
// 			MorphId: mId,
// 		}, metadataJson.Attributes, false, rarityResult)
// 		res, err := handlers.PersistSinglePolymorph(morphEntity, polymorphDBName, rarityCollectionName, mEvent.OldGene.String(), geneDifferences)
// 		if err != nil {
// 			log.Println(err)
// 		} else {
// 			log.Println(res)
// 		}

// 		if !hasTxMap {
// 			txMap = make(map[uint]bool)
// 			txState[morphEvent.TxHash.Hex()] = txMap
// 		}
// 		txState[morphEvent.TxHash.Hex()][morphEvent.Index] = true
// 		go handlers.SaveTransaction(polymorphDBName, transactionsCollectionName, models.Transaction{
// 			BlockNumber: morphEvent.BlockNumber,
// 			TxIndex:     morphEvent.TxIndex,
// 			TxHash:      morphEvent.TxHash.Hex(),
// 			LogIndex:    morphEvent.Index,
// 		})
// 	} else if txMap[morphEvent.Index] {
// 		log.Println("Already processed morph event! Skipping...")
// 	}
// }
