package services

import (
	"log"
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
)

func PollProcess(ethClient *dlt.EthereumClient, contractAbi abi.ABI, instance *store.Store, address string, configService *config.ConfigService, dbInfo rarityTypes.DBInfo, txState map[string]map[uint]bool) {
	var wg sync.WaitGroup
	eventLogsMutex := rarityTypes.EventLogsMutex{EventSigs: make(map[string]int), EventLogs: []types.Log{}}
	genesMap := make(map[string]string)
	tokenToMorphEvent := make(map[string]types.Log)

	lastProcessedBlockNumber := collectEvents(ethClient, contractAbi, instance, address, configService, dbInfo.PolymorphDBName, dbInfo.RarityCollectionName, dbInfo.BlocksCollectionName, 0, 0, &wg, &eventLogsMutex)

	helpers.SortMorphEvents(eventLogsMutex.EventLogs)

	for _, ethLog := range eventLogsMutex.EventLogs {
		eventSig := ethLog.Topics[0].String()
		switch eventSig {
		case config.TokenMorphedSignature:
			processMorphs(ethLog, &wg, contractAbi, instance, configService, dbInfo.PolymorphDBName, dbInfo.RarityCollectionName, dbInfo.TransactionsCollectionName, txState, genesMap, tokenToMorphEvent)
		}
	}

	handlers.UpdateAllRanking(dbInfo.PolymorphDBName, dbInfo.RarityCollectionName)
	res, err := handlers.CreateOrUpdateLastProcessedBlock(lastProcessedBlockNumber, dbInfo.PolymorphDBName, dbInfo.BlocksCollectionName)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(res)
	}
}

func processMorphs(morphEvent types.Log, wg *sync.WaitGroup, contractAbi abi.ABI, instance *store.Store, configService *config.ConfigService, polymorphDBName string,
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
		geneDifferences := helpers.DetectGeneDifferences(mEvent.OldGene.String(), mEvent.NewGene.String())

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
