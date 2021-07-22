package handlers

import (
	"context"
	"encoding/json"
	"log"
	"rarity-backend/db"
	"rarity-backend/rarityTypes"

	"go.mongodb.org/mongo-driver/bson"
)

func GetTransactionsMapping(polymorphDBName string, transactionsColl string) map[string]map[uint]bool {
	collection, err := db.GetMongoDbCollection(polymorphDBName, transactionsColl)
	if err != nil {
		log.Fatalln(err)
	}

	var transactions []rarityTypes.Transaction
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

func SaveTransaction(polymorphDBName string, transactionsColl string, transaction rarityTypes.Transaction) {
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

	log.Printf("Inserted new transaction in DB:\ntxHash: %v\nLogIndex: %v", transaction.TxHash, transaction.LogIndex)
	log.Println()
}
