package handlers

import (
	"context"
	"log"
	"rarity-backend/constants"
	"rarity-backend/db"
	"rarity-backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetMorphPriceMapping fetches all records from the morph cost collection. Returns a mapping of the records
//
// The application needs to track the changes in the morph prices in order to create correct morph history snapshots
func GetMorphPriceMapping(polymorphDBName string, priceCollection string) map[string]float32 {
	collection, err := db.GetMongoDbCollection(polymorphDBName, priceCollection)
	if err != nil {
		log.Fatalln(err)
	}

	var morphPrices []models.MorphCost
	results, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println(err)
	}
	results.All(context.Background(), &morphPrices)

	priceMap := make(map[string]float32)

	for _, price := range morphPrices {
		priceMap[price.TokenId] = price.Price
	}

	return priceMap
}

// SaveMorphPrice persists the new polymorph morph price to the database
//
// This price will be fetched and stored in memory every time the process starts.
func SaveMorphPrice(morphPrice models.MorphCost, polymorphDBName string, priceCollection string) {
	collection, err := db.GetMongoDbCollection(polymorphDBName, priceCollection)
	if err != nil {
		log.Fatalln(err)
	}

	update := bson.M{
		"$set": morphPrice,
	}

	opts := options.Update().SetUpsert(true)

	filter := bson.M{constants.MorphFieldNames.TokenId: morphPrice.TokenId}

	_, err = collection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("\nInserted new morph cost in DB:\n#:%v\nPrice: %v\n", morphPrice.TokenId, morphPrice.Price)
}
