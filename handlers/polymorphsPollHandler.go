package handlers

import (
	"context"
	"fmt"
	"log"
	"rarity-backend/constants"
	"rarity-backend/db"
	"rarity-backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PersistSinglePolymorph(entity models.PolymorphEntity, polymorphDBName string, rarityCollectionName string, oldGene string, geneDiff int) (string, error) {
	collection, err := db.GetMongoDbCollection(polymorphDBName, rarityCollectionName)
	if err != nil {
		return "", err
	}

	opts := options.Update().SetUpsert(true)
	filter := bson.M{constants.MorphFieldNames.TokenId: entity.TokenId}
	update := bson.M{}
	update["$set"] = entity

	if geneDiff > 0 && geneDiff <= 2 {
		update["$push"] = bson.M{constants.MorphFieldNames.OldGenes: oldGene}
		update["$inc"] = bson.M{constants.MorphFieldNames.Morphs: 1}
	} else if geneDiff > 2 {
		update["$push"] = bson.M{constants.MorphFieldNames.OldGenes: oldGene}
		update["$inc"] = bson.M{constants.MorphFieldNames.Scrambles: 1}
	}
	res, err := collection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return "", err
	}

	if res.UpsertedCount != 0 {
		return "Inserted id in polymorph db: " + entity.TokenId, nil
	} else if res.ModifiedCount != 0 {
		return "Updated id in polymorph db: " + entity.TokenId, nil
	} else {
		return "Didn't do shit in polymorph db (probably score is the same): " + entity.TokenId, nil
	}
}

func PersistMultiplePolymorphs(operations []mongo.WriteModel, polymorphDBName string, rarityCollectionName string) error {
	collection, err := db.GetMongoDbCollection(polymorphDBName, rarityCollectionName)
	if err != nil {
		return err
	}

	bulkOption := options.BulkWriteOptions{}

	res, err := collection.BulkWrite(context.Background(), operations, &bulkOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v entities' rank in polymorph db", res.ModifiedCount)
	return nil
}

func PersistMintEvents(bsonDocs []interface{}, polymorphDBName string, rarityCollectionName string) {
	collection, err := db.GetMongoDbCollection(polymorphDBName, rarityCollectionName)
	if err != nil {
		log.Fatal(err)
	}

	res, err := collection.InsertMany(context.Background(), bsonDocs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted %v polymorphs in DB", len(res.InsertedIDs))
}
