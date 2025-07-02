package types

type Token struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
	ChainId  int    `json:"chainId"`
}
