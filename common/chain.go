package common

import (
	"github.com/imduchuyyy/helix-wallet/evm"
	"github.com/imduchuyyy/helix-wallet/types"
)

var CHAIN = map[string]func(entropy string) types.Action{
	"eth": func(entropy string) types.Action {
		return evm.New(entropy, "Ethereum", "https://eth.llamarpc.com", "https://raw.githubusercontent.com/Uniswap/default-token-list/refs/heads/main/src/tokens/mainnet.json")
	},
}

func GetChainAction(chainName string, entropy string) (types.Action, bool) {
	genActionFunc, exists := CHAIN[chainName]
	if !exists {
		return nil, false
	}
	return genActionFunc(entropy), true
}
