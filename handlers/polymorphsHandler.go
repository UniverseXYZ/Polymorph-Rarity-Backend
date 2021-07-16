package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"rarity-backend/config"
	"rarity-backend/db"
	"rarity-backend/models"
	"strconv"
	"sync"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const RESULTS_LIMIT int64 = 10000

type UpdateModelMutex struct {
	mutex      sync.Mutex
	operations []mongo.WriteModel
}

// TODO: Create connection to db here and pass it to handlers
func GetPolymorphs(c *fiber.Ctx) {
	collection, err := db.GetMongoDbCollection(config.POLYMORPH_DB, config.RARITY_COLLECTION)
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
	collection, err := db.GetMongoDbCollection(config.POLYMORPH_DB, config.RARITY_COLLECTION)
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
	collection, err := db.GetMongoDbCollection(config.POLYMORPH_DB, config.RARITY_COLLECTION)
	if err != nil {
		return "", err
	}
	// This option will create new entity if no matching is found
	opts := options.Update().SetUpsert(true)
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

func CreateOrUpdatePolymorphEntities(entities []models.PolymorphEntity) error {
	collection, err := db.GetMongoDbCollection(config.POLYMORPH_DB, config.RARITY_COLLECTION)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	updateOperations := UpdateModelMutex{}
	for _, ent := range entities {
		wg.Add(1)
		go createWriteOperations(ent, &updateOperations, &wg)
	}
	wg.Wait()
	bulkOption := options.BulkWriteOptions{}

	res, err := collection.BulkWrite(context.Background(), updateOperations.operations, &bulkOption)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Updated %v entities' rank in polymorph db", res.ModifiedCount)
	return nil
}

func createWriteOperations(entity models.PolymorphEntity, mutex *UpdateModelMutex, wg *sync.WaitGroup) {
	operation := mongo.NewUpdateOneModel()
	mutex.mutex.Lock()
	operation.SetFilter(bson.M{"tokenid": entity.TokenId})
	operation.SetUpdate(bson.M{"$set": bson.M{"rank": entity.Rank}})
	mutex.operations = append(mutex.operations, operation)
	wg.Done()
	mutex.mutex.Unlock()

}
