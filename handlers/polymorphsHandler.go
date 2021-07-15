package handlers

import (
	"context"
	"encoding/json"
	"rarity-backend/db"
	"rarity-backend/models"
	"strconv"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const RESULTS_LIMIT int64 = 1000

// TODO: Create connection to db here and pass it to handlers
func GetPolymorphs(c *fiber.Ctx) {
	collection, err := db.GetMongoDbCollection("polymorphs-rarity", "rarity")
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	var findOptions options.FindOptions

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
	sortDir, err := strconv.Atoi(c.Query("sortDir"))
	if err != nil {
		sortDir = 1
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
	collection, err := db.GetMongoDbCollection("polymorphs-rarity", "rarity")
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

func CreateOrUpdatePolymorphEntity(entity models.PolymorphEntity) (string, error) {
	collection, err := db.GetMongoDbCollection("polymorphs-rarity", "rarity")
	if err != nil {
		return "", err
	}
	// This option will create new entity if no matching is found
	opts := options.Update().SetUpsert(true)
	// entity := models.PolymorphEntity{TokenId: event.MorphId.String(), Gene: event.NewGene.String(), RarityScore: rarityScore, IsVirgin: setVirgin}
	filter := bson.M{"tokenid": entity.TokenId}
	update := bson.M{
		"$set": entity,
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
