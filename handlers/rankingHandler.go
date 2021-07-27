package handlers

import (
	"context"
	"log"
	"rarity-backend/constants"
	"rarity-backend/db"
	"rarity-backend/models"
	"rarity-backend/structs"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateAllRanking(polymorphDBName string, rarityCollectionName string) {
	ranking := structs.RankMutex{}
	collection, err := db.GetMongoDbCollection(polymorphDBName, rarityCollectionName)
	if err != nil {
		log.Println(err)
	}

	var entities []models.PolymorphEntity

	var findOptions options.FindOptions
	findOptions.SetLimit(10000)
	findOptions.SetSort(bson.D{{constants.MorphFieldNames.RarityScore, -1}})
	results, err := collection.Find(context.Background(), bson.D{}, &findOptions)
	if err != nil {
		log.Println(err)
	}

	results.All(context.Background(), &entities)

	var wg sync.WaitGroup
	for _, entity := range entities {
		wg.Add(1)
		setRank(entity, &ranking, &wg)
	}
	wg.Wait()
	if len(ranking.Operations) > 0 {
		err = PersistMultiplePolymorphs(ranking.Operations, polymorphDBName, rarityCollectionName)
		if err != nil {
			log.Println(err)
		}
	}
}

func setRank(entity models.PolymorphEntity, ranking *structs.RankMutex, wg *sync.WaitGroup) {
	ranking.Mutex.Lock()

	if ranking.PrevRarity != entity.RarityScore {
		ranking.Rank++
		ranking.PrevRarity = entity.RarityScore
	}

	if entity.Rank != ranking.Rank {
		operation := mongo.NewUpdateOneModel()
		operation.SetFilter(bson.M{constants.MorphFieldNames.TokenId: entity.TokenId})
		operation.SetUpdate(bson.M{"$set": bson.M{constants.MorphFieldNames.Rank: ranking.Rank}})
		ranking.Operations = append(ranking.Operations, operation)
	}
	ranking.Mutex.Unlock()
	wg.Done()
}
