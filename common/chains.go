package common

import (
	"github.com/imduchuyyy/helix-wallet/evm"
	"github.com/imduchuyyy/helix-wallet/types"
)

var CHAIN = map[string]types.ChainConfig{
	"eth": {
		Chain: types.Chain{
			Name: "Ethereum",
			Rpcs: []string{
				"https://eth.llamarpc.com",
				"https://eth1.lava.build",
				"https://eth-mainnet.public.blastapi.io",
			},
			TokenListURL: "https://raw.githubusercontent.com/Uniswap/default-token-list/refs/heads/main/src/tokens/mainnet.json",
		},
		GenChainAction: func(entropy string) types.Action {
			return evm.New(entropy)
		},
	},
}

func GetChain(chainName string, entropy string) (types.Chain, bool) {
	config, exists := CHAIN[chainName]
	if !exists {
		return types.Chain{}, false
	}
	config.Chain.Action = config.GenChainAction(entropy)
	return config.Chain, true
}
