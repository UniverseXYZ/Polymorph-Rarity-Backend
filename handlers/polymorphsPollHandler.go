package handlers

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"os"
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
		log.Println("Error updating polymorph entity in rarity v2 collection.")
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
func PersistMintEvents(logs *[]types.Log, bsonDocs []interface{}, polymorphDBName string, rarityCollectionName string) error {
	collection, err := db.GetMongoDbCollection(polymorphDBName, rarityCollectionName)
	if err != nil {
		return err
	}
	res, err := collection.InsertMany(context.Background(), bsonDocs)
	if err != nil {
		log.Printf("Error inserting many documents in MongoDB %v", err)
		return err
	}
	log.Println(fmt.Sprintf("Inserted %v polymorphs in DB", len(res.InsertedIDs)))
	return nil
}

// DeleteV1Rarity Deletes all polymorph records from all V1 collections (after burnToMint, all info about v1 should disappear)
func DeleteV1Rarity(polymorphDBName string, newlyMinted *[]types.Log) error {

	raritiesV1 := os.Getenv("RARITIES_V1")
	historyV1 := os.Getenv("HISTORY_V1")
	morphCostV1 := os.Getenv("MORPH_COST_V1")

	collectionRaritiesV1, err := db.GetMongoDbCollection(polymorphDBName, raritiesV1)
	if err != nil {
		return err
	}

	historyV1Collection, err := db.GetMongoDbCollection(polymorphDBName, historyV1)
	if err != nil {
		return err
	}

	morphCostV1Collection, err := db.GetMongoDbCollection(polymorphDBName, morphCostV1)
	if err != nil {
		return err
	}

	// Delete ids of newly minted from V1 Collections
	for i := 0; i < len(*newlyMinted); i++ {
		currentPolymorphIdToDelete := (*newlyMinted)[i].Topics[1].Big().Int64()
		filter := bson.M{"tokenid": currentPolymorphIdToDelete}

		// Delete from rarities-v1 collection
		_, err = collectionRaritiesV1.DeleteOne(context.Background(), filter)
		if err != nil {
			log.Printf("Error deleting document from RaritiesV1 collection. %v ", err)
		} else {
			log.Println(fmt.Sprintf("Deleted polymorph #[%v] record from rarities-v1 collection", currentPolymorphIdToDelete))
		}

		// delete from history-v1 collection
		_, err = historyV1Collection.DeleteOne(context.Background(), filter)
		if err != nil {
			log.Printf("Error deleting document from HistoryV1 collection. %v ", err)
		} else {
			log.Println(fmt.Sprintf("Deleted polymorph #[%v] record from history-v1 collection", currentPolymorphIdToDelete))
		}

		// delete from morph-cost collection
		_, err = morphCostV1Collection.DeleteOne(context.Background(), filter)
		if err != nil {
			log.Printf("Error deleting document from MorphCostV1 collection. %v ", err)
		} else {
			log.Println(fmt.Sprintf("Deleted polymorph #[%v] record from morph-cost-v1", currentPolymorphIdToDelete))
		}
	}
	return err
}
