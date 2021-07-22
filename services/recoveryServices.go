package services

import (
	"encoding/json"
	"log"
	"math/big"
	"rarity-backend/config"
	"rarity-backend/dlt"
	"rarity-backend/handlers"
	"rarity-backend/helpers"
	"rarity-backend/metadata"
	"rarity-backend/rarityIndex"
	"rarity-backend/rarityTypes"
	"rarity-backend/store"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/bson"
)

func RecoverProcess(ethClient *dlt.EthereumClient, contractAbi abi.ABI, instance *store.Store, address string, configService *config.ConfigService, dbInfo rarityTypes.DBInfo, txState map[string]map[uint]bool) {
	var wg sync.WaitGroup
	mintsMutex := rarityTypes.MintsMutex{TokensMap: make(map[string]bool)}
	eventLogsMutex := rarityTypes.EventLogsMutex{EventSigs: make(map[string]int), EventLogs: []types.Log{}}
	genesMap := make(map[string]string)
	tokenToMorphEvent := make(map[string]types.Log)

	// Collect mint and morph events
	lastProcessedBlockNumber := collectEvents(ethClient, contractAbi, instance, address, configService, dbInfo.PolymorphDBName, dbInfo.RarityCollectionName, dbInfo.BlocksCollectionName, 0, 0, &wg, &eventLogsMutex)

	// Persist mints
	for _, ethLog := range eventLogsMutex.EventLogs {
		eventSig := ethLog.Topics[0].String()
		switch eventSig {
		case config.TokenMintedSignature:
			wg.Add(1)
			go processMint(ethLog, &wg, contractAbi, configService, dbInfo.PolymorphDBName, dbInfo.RarityCollectionName, &mintsMutex)
		}
	}
	wg.Wait()
	if len(mintsMutex.Documents) > 0 && len(mintsMutex.Documents) == 10000 {
		handlers.InsertManyMintEvents(mintsMutex.Documents, dbInfo.PolymorphDBName, dbInfo.RarityCollectionName)
	} else if len(mintsMutex.Documents) > 0 && len(mintsMutex.Documents) != 10000 {
		log.Printf("Expected 10000 mints, found %v. Please check config and restart app", len(mintsMutex.Documents))
		handlers.InsertManyMintEvents(mintsMutex.Documents, dbInfo.PolymorphDBName, dbInfo.RarityCollectionName)
	}

	// Sort polymorphs
	helpers.SortMorphEvents(eventLogsMutex.EventLogs)
	// Persist Morphs
	for _, ethLog := range eventLogsMutex.EventLogs {
		eventSig := ethLog.Topics[0].String()
		switch eventSig {
		case config.TokenMorphedSignature:
			processInitialMorphs(ethLog, &wg, contractAbi, instance, configService, dbInfo.PolymorphDBName, dbInfo.RarityCollectionName, dbInfo.TransactionsCollectionName, txState, genesMap, tokenToMorphEvent)
		}
	}

	// Persist leftover morphs
	for id := range genesMap {
		ethLog := tokenToMorphEvent[id]
		processLeftoverMorphs(ethLog, &wg, contractAbi, instance, configService, dbInfo.PolymorphDBName, dbInfo.RarityCollectionName, dbInfo.TransactionsCollectionName, txState, genesMap)
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

func processMint(mintEvent types.Log, wg *sync.WaitGroup, contractAbi abi.ABI, configService *config.ConfigService, polymorphDBName string, rarityCollectionName string, mintsMutex *rarityTypes.MintsMutex) {
	defer wg.Done()
	var event rarityTypes.PolymorphEvent
	mintsMutex.Mutex.Lock()
	contractAbi.UnpackIntoInterface(&event, "TokenMinted", mintEvent.Data)
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

func processInitialMorphs(morphEvent types.Log, wg *sync.WaitGroup, contractAbi abi.ABI, instance *store.Store, configService *config.ConfigService, polymorphDBName string,
	rarityCollectionName string, transactionsCollectionName string, txState map[string]map[uint]bool, oldGenesMap map[string]string, tokenToMorphEvent map[string]types.Log) {
	var mEvent rarityTypes.MorphedEvent
	err := contractAbi.UnpackIntoInterface(&mEvent, "TokenMorphed", morphEvent.Data)
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
		var geneDifferences int
		if oldGenesMap[mId.String()] != "" {
			geneDifferences = helpers.DetectGeneDifferences(oldGenesMap[mId.String()], mEvent.OldGene.String())
		}
		oldGenesMap[mId.String()] = mEvent.OldGene.String()
		tokenToMorphEvent[mId.String()] = morphEvent

		g := metadata.Genome(mEvent.NewGene.String())
		metadataJson := (&g).Metadata(mId.String(), configService)

		rarityResult := rarityIndex.CalulateRarityScore(metadataJson.Attributes, false)
		morphEntity := helpers.CreateMorphEntity(rarityTypes.PolymorphEvent{
			NewGene: mEvent.NewGene,
			OldGene: mEvent.OldGene,
			MorphId: mId,
		}, metadataJson.Attributes, false, rarityResult)
		res, err := handlers.CreateOrUpdatePolymorphEntity(morphEntity, polymorphDBName, rarityCollectionName, mEvent.OldGene.String(), geneDifferences)
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
		go handlers.SaveTransaction(polymorphDBName, transactionsCollectionName, rarityTypes.Transaction{
			BlockNumber: morphEvent.BlockNumber,
			TxIndex:     morphEvent.TxIndex,
			TxHash:      morphEvent.TxHash.Hex(),
			LogIndex:    morphEvent.Index,
		})
	} else if txMap[morphEvent.Index] {
		log.Println("Already processed morph event! Skipping...")
	}
}

func processLeftoverMorphs(morphEvent types.Log, wg *sync.WaitGroup, contractAbi abi.ABI, instance *store.Store, configService *config.ConfigService, polymorphDBName string,
	rarityCollectionName string, transactionsCollectionName string, txState map[string]map[uint]bool, oldGenesMap map[string]string) {
	var mEvent rarityTypes.MorphedEvent
	err := contractAbi.UnpackIntoInterface(&mEvent, "TokenMorphed", morphEvent.Data)
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

	geneDifferences := helpers.DetectGeneDifferences(oldGenesMap[mId.String()], mEvent.NewGene.String())

	g := metadata.Genome(mEvent.NewGene.String())
	metadataJson := (&g).Metadata(mId.String(), configService)

	rarityResult := rarityIndex.CalulateRarityScore(metadataJson.Attributes, false)
	morphEntity := helpers.CreateMorphEntity(rarityTypes.PolymorphEvent{
		NewGene: mEvent.NewGene,
		OldGene: mEvent.OldGene,
		MorphId: mId,
	}, metadataJson.Attributes, false, rarityResult)

	res, err := handlers.CreateOrUpdateLeftoverPolymorphEntity(morphEntity, polymorphDBName, rarityCollectionName, mEvent.OldGene.String(), geneDifferences)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(res)
	}
}
