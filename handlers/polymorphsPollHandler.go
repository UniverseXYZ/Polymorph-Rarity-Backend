package handlers

import (
	"context"
	"fmt"
	"log"
	"rarity-backend/constants"
	"rarity-backend/db"
	"rarity-backend/models"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PersistSinglePolymorph persists the polymorph entity in the rarities collection.
//
// Depending on the number of gene differences either the scramble or morph fields will be incremented.
// The old gene will also be appended to the oldGenes field. It's currently used to manually verify if the persisted entities and history snapshot are accurate
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
		return "Inserted id in polymorph db: " + strconv.Itoa(entity.TokenId), nil
	} else if res.ModifiedCount != 0 {
		return "Updated id in polymorph db: " + strconv.Itoa(entity.TokenId), nil
	} else {
		return "Didn't do shit in polymorph db (probably score is the same): " + strconv.Itoa(entity.TokenId), nil
	}
}

// PersistMultiplePolymorphs persists multiple changed polymorph entities in one go.
//
// Bulk writing to database saves a lot of time
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
	log.Println(fmt.Sprintf("Updated %v entities' rank in polymorph db", res.ModifiedCount))
	return nil
}

// PersistMintEvents persists all the processed mints in the database in one go.
//
// Bulk writing to database saves a lot of time
func PersistMintEvents(bsonDocs []interface{}, polymorphDBName string, rarityCollectionName string) {
	collection, err := db.GetMongoDbCollection(polymorphDBName, rarityCollectionName)
	if err != nil {
		log.Fatal(err)
	}

	res, err := collection.InsertMany(context.Background(), bsonDocs)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fmt.Sprintf("Inserted %v polymorphs in DB", len(res.InsertedIDs)))
}
