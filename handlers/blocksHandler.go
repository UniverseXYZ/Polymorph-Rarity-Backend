package handlers

import (
	"context"
	"rarity-backend/db"
	"rarity-backend/models"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetLastProcessedBlockNumber() (int64, error) {
	collection, err := db.GetMongoDbCollection("polymorphs-rarity", "processed-block")
	if err != nil {
		return 0, err
	}

	lastBlock := collection.FindOne(context.Background(), bson.M{})

	if lastBlock.Err() != nil {
		return 0, lastBlock.Err()
	}

	var result bson.M
	lastBlock.Decode(&result)

	if result == nil {
		return 0, err
	}

	// HOW TO GET IT
	lastProcessedBlockNumber := result["number"]
	block := lastProcessedBlockNumber.(int64)
	return block, nil
}

func CreateOrUpdateLastProcessedBlock(number uint64) (string, error) {
	collection, err := db.GetMongoDbCollection("polymorphs-rarity", "processed-block")
	if err != nil {
		return "", err
	}

	entity := models.ProcessedBlockEntity{Number: number}

	update := bson.M{
		"$set": entity,
	}

	// This option will create new entity if no matching is found
	opts := options.Update().SetUpsert(true)

	objID, _ := primitive.ObjectIDFromHex(strconv.FormatInt(0, 16))
	filter := bson.M{"_id": objID}

	_, err = collection.UpdateOne(context.Background(), filter, update, opts)

	if err != nil {
		return "", nil
	}

	return "Successfully persisted new last processed block number: " + strconv.FormatUint(number, 10), nil
}
