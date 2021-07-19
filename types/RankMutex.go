package types

import (
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

type RankMutex struct {
	Rank       int
	PrevRarity int
	Operations []mongo.WriteModel
	Mutex      sync.Mutex
}
