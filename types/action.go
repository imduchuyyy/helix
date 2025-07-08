package types

type Action interface {
	GetAddress() (string, error)
	GetPrivateKey() (string, error)
	GetTokenBalances() ([]Token, error)
	ChainName() string
}
