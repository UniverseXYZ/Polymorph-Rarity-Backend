package handlers

import (
	"context"
	"log"
	"rarity-backend/db"
	"rarity-backend/models"
	"rarity-backend/types"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateRanking(currEntity models.PolymorphEntity, polymorphDBName string, rarityCollectionName string) {
	ranking := types.RankMutex{}
	collection, err := db.GetMongoDbCollection(polymorphDBName, rarityCollectionName)
	if err != nil {
		log.Println(err)
	}

	var entities []models.PolymorphEntity
	results, err := collection.Find(context.Background(), bson.M{"rarirtyscore": bson.M{"$lte": currEntity.RarityScore}})
	if err != nil {
		log.Println(err)
	}

	results.All(context.Background(), &entities)

	if len(entities) == 0 {
		var findOptions options.FindOptions
		findOptions.SetLimit(10000)
		findOptions.SetSort(bson.D{{"rarityscore", -1}})
		results, err = collection.Find(context.Background(), bson.D{}, &findOptions)
		if err != nil {
			log.Println(err)
		}

		results.All(context.Background(), &entities)
	}

	var wg sync.WaitGroup
	for _, entity := range entities {
		wg.Add(1)
		setRank(entity, &ranking, &wg)
	}
	wg.Wait()

	err = CreateOrUpdatePolymorphEntities(ranking.Operations, polymorphDBName, rarityCollectionName)
	if err != nil {
		log.Println(err)
	}
}

func setRank(entity models.PolymorphEntity, ranking *types.RankMutex, wg *sync.WaitGroup) {
	ranking.Mutex.Lock()

	if ranking.PrevRarity != entity.RarityScore {
		ranking.Rank++
		ranking.PrevRarity = entity.RarityScore
	}

	if entity.Rank != ranking.Rank {
		operation := mongo.NewUpdateOneModel()
		operation.SetFilter(bson.M{"tokenid": entity.TokenId})
		operation.SetUpdate(bson.M{"$set": bson.M{"rank": ranking.Rank}})
		ranking.Operations = append(ranking.Operations, operation)
	}
	ranking.Mutex.Unlock()
	wg.Done()
}
