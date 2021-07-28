package services

import (
	"context"
	"encoding/json"
	"log"
	"math/big"
	"rarity-backend/constants"
	"rarity-backend/dlt"
	"rarity-backend/handlers"
	"rarity-backend/helpers"
	"rarity-backend/metadata"
	"rarity-backend/models"
	"rarity-backend/rarityIndex"
	"rarity-backend/store"
	"rarity-backend/structs"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/bson"
)

func RecoverProcess(ethClient *dlt.EthereumClient, contractAbi abi.ABI, instance *store.Store, address string, configService *structs.ConfigService,
	dbInfo structs.DBInfo, txState map[string]map[uint]bool, morphCostMap map[string]float32) {
	var wg sync.WaitGroup
	mintsMutex := structs.MintsMutex{TokensMap: make(map[string]bool)}
	eventLogsMutex := structs.EventLogsMutex{EventSigs: make(map[string]int), EventLogs: []types.Log{}}
	genesMap := make(map[string]string)
	tokenToMorphEvent := make(map[string]types.Log)

	// Collect mint and morph events
	lastProcessedBlockNumber := collectEvents(ethClient, contractAbi, instance, address, configService, dbInfo.PolymorphDBName, dbInfo.RarityCollectionName, dbInfo.BlocksCollectionName, 0, 0, &wg, &eventLogsMutex)

	// Persist mints
	for _, ethLog := range eventLogsMutex.EventLogs {
		eventSig := ethLog.Topics[0].String()
		switch eventSig {
		case constants.MintEvent.Signature:
			wg.Add(1)
			go processMint(ethLog, &wg, contractAbi, configService, dbInfo.PolymorphDBName, dbInfo.RarityCollectionName, &mintsMutex)
		}
	}

	wg.Wait()
	if len(mintsMutex.Documents) > 0 {
		handlers.PersistMintEvents(mintsMutex.Documents, dbInfo.PolymorphDBName, dbInfo.RarityCollectionName)
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
		processLeftoverMorphs(ethLog, ethClient, contractAbi, instance, configService, dbInfo, txState, genesMap, morphCostMap)
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
		rarityResult := rarityIndex.CalulateRarityScore(metadataJson.Attributes, true)
		mintEntity := helpers.CreateMorphEntity(event, metadataJson.Attributes, true, rarityResult)

		mintsMutex.Mints = append(mintsMutex.Mints, mintEntity)
		mintsMutex.TokensMap[event.MorphId.String()] = true
		var bdoc interface{}
		json, _ := json.Marshal(mintEntity)
		bson.UnmarshalExtJSON(json, false, &bdoc)
		mintsMutex.Documents = append(mintsMutex.Documents, bdoc)
	} else {
		log.Println("Empty gene mint event for morph id: " + event.MorphId.String())
	}
	mintsMutex.Mutex.Unlock()
}

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

		rarityResult := rarityIndex.CalulateRarityScore(metadataJson.Attributes, false)
		morphEntity := helpers.CreateMorphEntity(structs.PolymorphEvent{NewGene: mEvent.NewGene, OldGene: mEvent.OldGene, MorphId: mId}, metadataJson.Attributes, false, rarityResult)

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

func processLeftoverMorphs(morphEvent types.Log, ethClient *dlt.EthereumClient, contractAbi abi.ABI, instance *store.Store, configService *structs.ConfigService, dbInfo structs.DBInfo,
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

	g := metadata.Genome(mEvent.NewGene.String())
	metadataJson := (&g).Metadata(mId.String(), configService)

	rarityResult := rarityIndex.CalulateRarityScore(metadataJson.Attributes, false)
	morphEntity := helpers.CreateMorphEntity(structs.PolymorphEvent{NewGene: mEvent.NewGene, MorphId: mId}, metadataJson.Attributes, false, rarityResult)

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
