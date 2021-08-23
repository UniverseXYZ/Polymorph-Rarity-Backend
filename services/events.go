package services

import (
	"context"
	"log"
	"math/big"
	"rarity-backend/constants"
	"rarity-backend/dlt"
	"rarity-backend/handlers"
	"rarity-backend/store"
	"rarity-backend/structs"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// collectEvents sends request using the ethereum client for events emitted from the Polymorphs contract. It iterates over the events and filters for mint and morph events.
// Events iteration is implemented concurrently.
// Returns last processed block so it can be persisted in the database after the events have been fully processed.
//
// If events in the block range is > 10,000 the range is split in two and the function is called recursively until the blocks range can be processed.(10,000 limit: https://infura.io/docs/ethereum/json-rpc/eth_getLogs)
func collectEvents(ethClient *dlt.EthereumClient, contractAbi abi.ABI, instance *store.Store, address string, configService *structs.ConfigService, polymorphDBName string, rarityCollectionName string, blocksCollectionName string, startBlock int64, endBlock int64, wg *sync.WaitGroup, elm *structs.EventLogsMutex) uint64 {
	var lastProcessedBlockNumber, lastChainBlockNumberInt64 int64

	if startBlock != 0 {
		lastProcessedBlockNumber = startBlock
	} else {
		lastProcessedBlockNumber, _ = handlers.GetLastProcessedBlockNumber(polymorphDBName, blocksCollectionName)
	}

	if endBlock != 0 {
		lastChainBlockNumberInt64 = endBlock
	} else {
		lastChainBlockHeader, err := ethClient.Client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			log.Println(err)
			lastChainBlockNumberInt64 = int64(lastProcessedBlockNumber)
		}
		lastChainBlockNumberInt64 = int64(lastChainBlockHeader.Number.Uint64())

	}

	ethLogs, err := ethClient.Client.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(lastProcessedBlockNumber),
		ToBlock:   big.NewInt(lastChainBlockNumberInt64),
		Addresses: []common.Address{common.HexToAddress(address)},
	})
	if err != nil {
		log.Println(err)
		middle := (lastProcessedBlockNumber + lastChainBlockNumberInt64) / 2
		collectEvents(ethClient, contractAbi, instance, address, configService, polymorphDBName, rarityCollectionName, blocksCollectionName, lastProcessedBlockNumber, middle, wg, elm)
		collectEvents(ethClient, contractAbi, instance, address, configService, polymorphDBName, rarityCollectionName, blocksCollectionName, middle+1, lastChainBlockNumberInt64, wg, elm)
	} else {
		log.Printf("Processing blocks %v - %v for polymorph events", lastProcessedBlockNumber, lastChainBlockNumberInt64)
		wg.Add(1)
		go saveToEventLogMutex(ethLogs, elm, wg)
	}
	wg.Wait()
	return uint64(lastChainBlockNumberInt64)
}

// saveToEventLogMutex concurrently saves mint and morph events an array which will be processed after all events have been filtered for these events.
//
// Uses Mutex and WaitGroup to prevent race conditions
func saveToEventLogMutex(ethLogs []types.Log, elm *structs.EventLogsMutex, wg *sync.WaitGroup) {
	defer wg.Done()
	elm.Mutex.Lock()
	for _, ethLog := range ethLogs {
		eventSig := ethLog.Topics[0].String()
		switch eventSig {
		case constants.MintEvent.Signature, constants.MorphEvent.Signature:
			elm.EventLogs = append(elm.EventLogs, ethLog)
		}
	}
	elm.Mutex.Unlock()
}
