package handlers

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"rarity-backend/config"
	"rarity-backend/constants"
	"rarity-backend/db"
	"rarity-backend/helpers"
	"rarity-backend/structs"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
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
	var filters, searchFilters, aggrFilters = bson.M{}, bson.M{}, bson.M{}

	if queryParams.Search != "" {
		searchFilters = helpers.ParseSearchQueryString(queryParams.Search)
		for k, v := range searchFilters {
			aggrFilters[k] = v
		}

	}

	if queryParams.Filter != "" {
		filters = helpers.ParseFilterQueryString(queryParams.Filter)
		for k, v := range filters {
			aggrFilters[k] = v
		}
	}

	var findOptions options.FindOptions

	removePrivateFields(&findOptions)

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

	curr, err := collection.Find(context.Background(), aggrFilters, &findOptions)
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

	findOptions := options.FindOneOptions{}
	removePrivateFieldsSingle(&findOptions)

	var filter bson.M = bson.M{}
	if c.Params("id") != "" {
		id := c.Params("id")
		filter = bson.M{constants.MorphFieldNames.TokenId: id}
	}

	var result bson.M
	curr := collection.FindOne(context.Background(), filter, &findOptions)

	curr.Decode(&result)

	if result == nil {
		c.Send(result)
		return
	}

	json, _ := json.Marshal(result)
	c.Send(json)
}

func removePrivateFields(findOptions *options.FindOptions) {
	for _, field := range config.NO_PROJECTION_FIELDS {
		findOptions.SetProjection(bson.M{field: 0})
	}
}

func removePrivateFieldsSingle(findOptions *options.FindOneOptions) {
	for _, field := range config.NO_PROJECTION_FIELDS {
		findOptions.SetProjection(bson.M{field: 0})
	}
}
