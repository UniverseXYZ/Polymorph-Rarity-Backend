package types

import "math/big"

type TokenMorphedEvent struct {
	OldGene   *big.Int
	NewGene   *big.Int
	Price     *big.Int
	EventType *big.Int
	MorphId   *big.Int
}
