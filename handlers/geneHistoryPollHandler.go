package handlers

import (
	"context"
	"encoding/json"
	"log"
	"rarity-backend/db"
	"rarity-backend/models"

	"go.mongodb.org/mongo-driver/bson"
)

func SavePolymorphHistory(entity models.PolymorphHistory, polymorphDBName string, historyCollectionName string) {
	collection, err := db.GetMongoDbCollection(polymorphDBName, historyCollectionName)

	if err != nil {
		log.Println(err)
	}

	var bdoc interface{}
	json, _ := json.Marshal(entity)
	bson.UnmarshalExtJSON(json, false, &bdoc)

	_, err = collection.InsertOne(context.Background(), bdoc)
	if err != nil {
		log.Println(err)
	}

	log.Println("Inserted history snapshot for polymorph #" + entity.TokenId)
}
