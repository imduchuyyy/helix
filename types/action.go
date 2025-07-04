package types

type Action interface {
	GetAddress() (string, error)
	ChainName() string
}
