package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"rarity-backend/db"
	"rarity-backend/models"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const RESULTS_LIMIT int64 = 10000

func GetPolymorphs(c *fiber.Ctx) {
	godotenv.Load()
	polymorphDBName := os.Getenv("POLYMORPH_DB")
	rarityCollectionName := os.Getenv("RARITY_COLLECTION")
	collection, err := db.GetMongoDbCollection(polymorphDBName, rarityCollectionName)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	var findOptions options.FindOptions
	// TODO: Add select query param for the rest of the params
	findOptions.SetProjection(bson.M{"_id": 0})
	take, err := strconv.ParseInt(c.Query("take"), 10, 64)
	if err != nil || take > RESULTS_LIMIT {
		take = RESULTS_LIMIT
	}
	findOptions.SetLimit(take)

	page, err := strconv.ParseInt(c.Query("page"), 10, 64)
	if err != nil {
		page = 1
	}

	findOptions.SetSkip((page - 1) * take)

	sortField := c.Query("sortField")
	sortDirQuery := c.Query("sortDir") // desc, asc
	sortDir := 1

	if sortDirQuery == "desc" {
		sortDir = -1
	}

	if sortField != "" {
		findOptions.SetSort(bson.D{{sortField, sortDir}})
	}

	curr, err := collection.Find(context.Background(), bson.D{}, &findOptions)
	if err != nil {
		c.Status(500).Send(err)
	}

	defer curr.Close(context.Background())

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	var results []bson.M
	curr.All(context.Background(), &results)

	if results == nil {
		c.SendStatus(404)
		return
	}

	json, _ := json.Marshal(results)
	c.Send(json)
}

func GetPolymorphById(c *fiber.Ctx) {
	godotenv.Load()

	polymorphDBName := os.Getenv("POLYMORPH_DB")
	rarityCollectionName := os.Getenv("RARITY_COLLECTION")

	collection, err := db.GetMongoDbCollection(polymorphDBName, rarityCollectionName)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	var filter bson.M = bson.M{}
	if c.Params("id") != "" {
		id := c.Params("id")
		filter = bson.M{"tokenid": id}
	}

	var result bson.M
	curr := collection.FindOne(context.Background(), filter)

	curr.Decode(&result)

	if result == nil {
		c.SendStatus(404)
		return
	}

	json, _ := json.Marshal(result)
	c.Send(json)
}

func CreateOrUpdatePolymorphEntity(entity models.PolymorphEntity, polymorphDBName string, rarityCollectionName string, oldGene string, geneDiff int) (string, error) {
	collection, err := db.GetMongoDbCollection(polymorphDBName, rarityCollectionName)
	if err != nil {
		return "", err
	}

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"tokenid": entity.TokenId}
	update := bson.M{}
	update["$set"] = entity

	if geneDiff > 0 && geneDiff <= 2 {
		update["$push"] = bson.M{"oldgenes": oldGene}
		update["$inc"] = bson.M{"morphs": 1}
	} else if geneDiff > 2 {
		update["$push"] = bson.M{"oldgenes": oldGene}
		update["$inc"] = bson.M{"scrambles": 1}
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

func CreateOrUpdatePolymorphEntities(operations []mongo.WriteModel, polymorphDBName string, rarityCollectionName string) error {
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

func InsertManyMintEvents(bsonDocs []interface{}, polymorphDBName string, rarityCollectionName string) {
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
