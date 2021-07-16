package services

import (
	"context"
	"log"
	"math/big"
	"rarity-backend/config"
	"rarity-backend/dlt"
	"rarity-backend/handlers"
	"rarity-backend/metadata"
	"rarity-backend/models"
	"rarity-backend/rarityIndex"
	"rarity-backend/store"
	"rarity-backend/types"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func ProcessBlocks(ethClient *dlt.EthereumClient, contractAbi abi.ABI, instance *store.Store, address string, configService *config.ConfigService, polymorphDBName string, rarityCollectionName string, blocksCollectionName string, startBlock int64, endBlock int64) {
	log.Println("Processing new blocks for morph events")

	var lastProcessedBlockNumber, lastChainBlockNumberInt64 int64
	if startBlock != 0 && endBlock != 0 {
		lastProcessedBlockNumber = startBlock
		lastChainBlockNumberInt64 = endBlock
	} else {
		lastProcessedBlockNumber, _ = handlers.GetLastProcessedBlockNumber(polymorphDBName, blocksCollectionName)
		lastChainBlockHeader, err := ethClient.Client.HeaderByNumber(context.Background(), nil)
		lastChainBlockNumberInt64 = int64(lastChainBlockHeader.Number.Uint64())

		if err != nil {
			log.Fatal(err)
			return
		}
	}

	ethLogs, err := ethClient.Client.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(lastProcessedBlockNumber),
		ToBlock:   big.NewInt(lastChainBlockNumberInt64),
		Addresses: []common.Address{common.HexToAddress(address)},
	})

	if err != nil {
		log.Println(err)
		middle := (lastProcessedBlockNumber + lastChainBlockNumberInt64) / 2
		ProcessBlocks(ethClient, contractAbi, instance, address, configService, polymorphDBName, rarityCollectionName, blocksCollectionName, lastProcessedBlockNumber, middle)
		ProcessBlocks(ethClient, contractAbi, instance, address, configService, polymorphDBName, rarityCollectionName, blocksCollectionName, middle+1, lastChainBlockNumberInt64)
	} else {

		tokenMintedSignature := "0x8c0bdd7bca83c4e0c810cbecf44bc544a9dc0b9f265664e31ce0ce85f07a052b"
		tokenMorphedSignature := "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"

		for _, vLog := range ethLogs {
			eventSig := vLog.Topics[0].String()
			switch eventSig {
			case tokenMintedSignature:
				processTokenMintedEvent(contractAbi, vLog.Data, vLog.Topics, configService, polymorphDBName, rarityCollectionName)
			case tokenMorphedSignature:
				processTokenMorphedEvent(vLog.Topics, configService, instance, polymorphDBName, rarityCollectionName)
			}
			//TODO: Add this as deferred somehow in order to save the last processed block number if app panicks
			persistProcessedBlock(vLog.BlockNumber, polymorphDBName, blocksCollectionName)
		}
	}
}

func persistProcessedBlock(blockNumber uint64, polymorphDBName string, blocksCollectionName string) {
	res, err := handlers.CreateOrUpdateLastProcessedBlock(blockNumber, polymorphDBName, blocksCollectionName)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(res)
	}

}

func processTokenMintedEvent(contractAbi abi.ABI, data []byte, topics []common.Hash, configService *config.ConfigService, polymorphDBName string, rarityCollectionName string) {
	var morphEvent types.PolymorphEvent

	err := contractAbi.UnpackIntoInterface(&morphEvent, "TokenMinted", data)
	if err != nil {
		log.Println(err)
		return
	}
	morphEvent.MorphId = topics[1].Big()
	morphEvent.OldGene = big.NewInt(0)
	if morphEvent.NewGene.String() == "0" {
		log.Println("Empty gene mint event for morph id: " + morphEvent.MorphId.String())
		return
	}
	processMorphAndPersist(morphEvent, configService, polymorphDBName, rarityCollectionName, true)
}

func processTokenMorphedEvent(topics []common.Hash, configService *config.ConfigService, contract *store.Store, polymorphDBName string, rarityCollectionName string) {
	morphEvent := types.PolymorphEvent{
		OldGene: topics[1].Big(),
		NewGene: topics[2].Big(),
		MorphId: topics[3].Big(),
	}
	result, err := contract.GeneOf(&bind.CallOpts{}, morphEvent.MorphId)
	if err != nil {
		log.Fatalln(err)
	}
	morphEvent.NewGene = result
	if morphEvent.OldGene.String() != "0" {
		processMorphAndPersist(morphEvent, configService, polymorphDBName, rarityCollectionName, false)
	}
}

func processMorphAndPersist(event types.PolymorphEvent, configService *config.ConfigService, polymorphDBName string, rarityCollectionName string, isVirgin bool) {
	g := metadata.Genome(event.NewGene.String())
	metadataJson := (&g).Metadata(event.MorphId.String(), configService)

	setName, hasCompletedSet, scaledRarity, matchingTraits, setMatchingTraits, colorMismatches := rarityIndex.CalulateRarityScore(metadataJson.Attributes, isVirgin)
	morphEntity := createMorphEntity(event, metadataJson.Attributes, setName, hasCompletedSet, isVirgin, scaledRarity, matchingTraits, setMatchingTraits, colorMismatches)
	res, err := handlers.CreateOrUpdatePolymorphEntity(morphEntity, polymorphDBName, rarityCollectionName)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(res)
	}
	handlers.UpdateRanking(morphEntity, polymorphDBName, rarityCollectionName)

}

func createMorphEntity(event types.PolymorphEvent, attributes []metadata.Attribute, setName string, hasCompletedSet bool, isVirgin bool, scaledRarity int, matchingTraits float64, setMatchingTraits []string, colorMismatches float64) models.PolymorphEntity {
	var background, leftHand, rightHand, head, eye, torso, pants, feet, character metadata.Attribute

	for _, attr := range attributes {
		switch attr.TraitType {
		case "Background":
			background = attr
		case "Character":
			character = attr
		case "Right Hand":
			rightHand = attr
		case "Left Hand":
			leftHand = attr
		case "Footwear":
			feet = attr
		case "Pants":
			pants = attr
		case "Torso":
			torso = attr
		case "Eyewear":
			eye = attr
		case "Headwear":
			head = attr
		}
	}

	morphEntity := models.PolymorphEntity{
		TokenId:         event.MorphId.String(),
		OldGene:         event.OldGene.String(),
		NewGene:         event.NewGene.String(),
		Headwear:        head.Value,
		Eyewear:         eye.Value,
		Torso:           torso.Value,
		Pants:           pants.Value,
		Footwear:        feet.Value,
		LeftHand:        leftHand.Value,
		RightHand:       rightHand.Value,
		Character:       character.Value,
		Background:      background.Value,
		RarityScore:     scaledRarity,
		IsVirgin:        isVirgin,
		MatchingTraits:  setMatchingTraits,
		ColorMismatches: int(colorMismatches),
	}
	return morphEntity
}
