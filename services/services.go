package services

import (
	"context"
	"log"
	"math/big"
	"rarity-backend/config"
	"rarity-backend/dlt"
	"rarity-backend/handlers"
	"rarity-backend/metadata"
	"rarity-backend/rarityIndex"
	"rarity-backend/store"
	"rarity-backend/types"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func ProcessBlocks(ethClient *dlt.EthereumClient, address string, configService *config.ConfigService) {
	log.Println("Processing new blocks for morph events")
	lastProcessedBlockNumber, err := handlers.GetLastProcessedBlockNumber()
	// If error, the recovery will start from block 0
	if err != nil {
		log.Println(err)
	}
	// TODO: Check if this really sends you the latest block or connection needs to be reset
	lastChainBlockHeader, err := ethClient.Client.HeaderByNumber(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
		return
	}

	lastChainBlockNumberInt64 := int64(lastChainBlockHeader.Number.Uint64())

	ethLogs, err := ethClient.Client.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(lastProcessedBlockNumber),
		ToBlock:   big.NewInt(lastChainBlockNumberInt64),
		Addresses: []common.Address{common.HexToAddress(address)},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	if err != nil {
		log.Fatal(err)
		return
	}

	tokenMintedSignature := "0x8c0bdd7bca83c4e0c810cbecf44bc544a9dc0b9f265664e31ce0ce85f07a052b"
	tokenMorphedSignature := "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	//	transferEventSignature := "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	// 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d
	// 0x5f7666687319b40936f33c188908d86aea154abd3f4127b4fa0a3f04f303c7da -- maybe is TokenMorphedEvent signature

	// eventSignatures := make(map[string]int)
	// for _, vLog := range ethLogs {
	// 	eventSig := vLog.Topics[0].String()
	// 	eventSignatures[eventSig]++
	// }

	for _, vLog := range ethLogs {
		eventSig := vLog.Topics[0].String()
		switch eventSig {
		case tokenMintedSignature:
			processTokenMintedEvent(contractAbi, vLog.Data, vLog.Topics, configService)
			//TODO: Add this as deferred somehow in order to save the last processed block number if app panicks
		case tokenMorphedSignature:
			processTokenMorphedEvent(contractAbi, vLog.Topics, configService)
			//TODO: Add this as deferred somehow in order to save the last processed block number if app panicks
		}

		res, err := handlers.CreateOrUpdateLastProcessedBlock(vLog.BlockNumber)
		if err != nil {
			log.Println(err)
		} else {
			log.Println(res)
		}
	}

}

func processTokenMintedEvent(contractAbi abi.ABI, data []byte, topics []common.Hash, configService *config.ConfigService) {
	var tokenMintedEvent types.TokenMintedEvent

	err := contractAbi.UnpackIntoInterface(&tokenMintedEvent, "TokenMinted", data)
	if err != nil {
		log.Println(err)
		return
	}
	tokenMintedEvent.MorphId = topics[1].Big()
	if tokenMintedEvent.NewGene.String() == "0" {
		log.Println("Empty gene mint event for morph id: " + tokenMintedEvent.MorphId.String())
		return
	}
	processMorphAndPersist(tokenMintedEvent, configService, true)
}

func processTokenMorphedEvent(contractAbi abi.ABI, topics []common.Hash, configService *config.ConfigService) {
	oldGene := topics[1].Big().String()
	event := types.TokenMintedEvent{NewGene: topics[2].Big(), MorphId: topics[3].Big()}

	// Morph event is emited after mint event so no need to write to db the same info
	if oldGene != "0" {
		processMorphAndPersist(event, configService, false)
	}
}

func processMorphAndPersist(event types.TokenMintedEvent, configService *config.ConfigService, isVirgin bool) {
	g := metadata.Genome(event.NewGene.String())
	metadataJson := (&g).Metadata(event.MorphId.String(), configService)

	rarityScore := rarityIndex.CalulateRarityScore(metadataJson.Attributes, isVirgin)

	res, err := handlers.CreateOrUpdatePolymorphEntity(event, rarityScore, isVirgin)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(res)
	}
}
