package handlers

import (
	"context"
	"log"
	"rarity-backend/config"
	"rarity-backend/db"
	"rarity-backend/models"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RankMutex struct {
	rank       int
	prevRarity int
	entities   []models.PolymorphEntity
	mutex      sync.Mutex
}

func UpdateRanking(currEntity models.PolymorphEntity) {
	ranking := RankMutex{}
	collection, err := db.GetMongoDbCollection(config.POLYMORPH_DB, config.RARITY_COLLECTION)
	if err != nil {
		log.Println(err)
	}

	entities := make([]models.PolymorphEntity, 10)
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

	err = CreateOrUpdatePolymorphEntities(ranking.entities)
	if err != nil {
		log.Println(err)
	}
}

func setRank(entity models.PolymorphEntity, ranking *RankMutex, wg *sync.WaitGroup) {
	ranking.mutex.Lock()
	if ranking.prevRarity != entity.RarityScore {
		ranking.rank++
		ranking.prevRarity = entity.RarityScore
	}
	entity.Rank = ranking.rank
	ranking.entities = append(ranking.entities, entity)

	ranking.mutex.Unlock()
	wg.Done()
}
