package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"rarity-backend/config"
	"rarity-backend/db"
	"rarity-backend/helpers"
	"rarity-backend/models"
	"rarity-backend/structs"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetPolymorphs(c *fiber.Ctx) {
	godotenv.Load()
	polymorphDBName := os.Getenv("POLYMORPH_DB")
	rarityCollectionName := os.Getenv("RARITY_COLLECTION")
	collection, err := db.GetMongoDbCollection(polymorphDBName, rarityCollectionName)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	queryParams := structs.QueryParams{}
	if err := c.QueryParser(&queryParams); err != nil {
		log.Println(err)
	}
	filters := bson.M{}
	searchFilters := bson.M{}

	if queryParams.Search != "" {
		searchFilters = helpers.ParseSearchQueryString(queryParams.Search)
	}

	if queryParams.Filter != "" {
		filters = helpers.ParseFilterQueryString(queryParams.Filter)
	}
	finalFilters := bson.M{}
	for k, v := range searchFilters {
		finalFilters[k] = v
	}
	for k, v := range filters {
		finalFilters[k] = v
	}
	var findOptions options.FindOptions
	// TODO: Add select query param for the rest of the params
	// for _, field := range strings.Split(filter.Select, ",") {

	// }
	findOptions.SetProjection(bson.M{"_id": 0})
	take, err := strconv.ParseInt(queryParams.Take, 10, 64)
	if err != nil || take > config.RESULTS_LIMIT {
		take = config.RESULTS_LIMIT
	}
	findOptions.SetLimit(take)

	page, err := strconv.ParseInt(queryParams.Page, 10, 64)
	if err != nil {
		page = 1
	}

	findOptions.SetSkip((page - 1) * take)

	sortDir := 1

	if queryParams.SortDir == "desc" {
		sortDir = -1
	}

	if queryParams.SortField != "" {
		findOptions.SetSort(bson.D{{queryParams.SortField, sortDir}})
	}

	curr, err := collection.Find(context.Background(), finalFilters, &findOptions)
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
		c.Send(results)
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
