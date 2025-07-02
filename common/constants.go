package common

import "github.com/imduchuyyy/helix-wallet/types"

var CHAIN = map[string]types.Chain{
	"eth": {
		Name: "Ethereum",
		Rpcs: []string{
			"https://eth.llamarpc.com",
			"https://eth1.lava.build",
			"https://eth-mainnet.public.blastapi.io",
		},
		TokenListURL: "https://raw.githubusercontent.com/Uniswap/default-token-list/refs/heads/main/src/tokens/mainnet.json",
	},
}
