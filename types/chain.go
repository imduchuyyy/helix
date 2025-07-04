package types

type ChainConfig struct {
	Chain          Chain
	GenChainAction func(entropy string) Action
}

type Chain struct {
	Name         string
	Rpcs         []string
	TokenListURL string
	Action       Action
}
