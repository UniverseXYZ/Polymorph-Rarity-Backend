package rarityTypes

import (
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
)

type EventLogsMutex struct {
	Mutex     sync.Mutex
	EventLogs []types.Log
	EventSigs map[string]int
}
