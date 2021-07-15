package types

import "math/big"

type TokenMintedEvent struct {
	NewGene *big.Int
	MorphId *big.Int
}
