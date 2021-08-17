package handlers

import (
	"context"
	"encoding/json"
	"log"
	"rarity-backend/db"
	"rarity-backend/models"

	"go.mongodb.org/mongo-driver/bson"
)

// GetTransactionsMapping fetches all records from the transactions collections. Returns a mapping of the records.
//
// The application has to know which morph events have already been processed in order to prevent duplicate false information stored in database
func GetTransactionsMapping(polymorphDBName string, transactionsColl string) map[string]map[uint]bool {
	collection, err := db.GetMongoDbCollection(polymorphDBName, transactionsColl)
	if err != nil {
		log.Fatalln(err)
	}

	var transactions []models.Transaction
	results, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println(err)
	}
	results.All(context.Background(), &transactions)

	txMap := make(map[string]map[uint]bool)

	for _, tx := range transactions {
		nestedMap, ok := txMap[tx.TxHash]
		if !ok {
			nestedMap = make(map[uint]bool)
			txMap[tx.TxHash] = nestedMap
		}
		txMap[tx.TxHash][tx.LogIndex] = true
	}

	return txMap
}

// SaveTransaction persists the processed transaction in the database
//
// If the application stops it will be able to load the processed event in memory from the database
func SaveTransaction(polymorphDBName string, transactionsColl string, transaction models.Transaction) {
	collection, err := db.GetMongoDbCollection(polymorphDBName, transactionsColl)
	if err != nil {
		log.Fatalln(err)
	}

	var bdoc interface{}
	json, _ := json.Marshal(transaction)
	bson.UnmarshalExtJSON(json, false, &bdoc)

	_, err = collection.InsertOne(context.Background(), bdoc)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("\nInserted new transaction in DB:\ntxHash: %v\nLogIndex: %v\n", transaction.TxHash, transaction.LogIndex)
}
