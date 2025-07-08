package types

import "math/big"

type Token struct {
	Name     string
	Address  string
	Symbol   string
	Decimals int
	Balance  *big.Int
}
