package common

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/imduchuyyy/helix-wallet/types"
)

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

var ZERO_ADDRESS = common.HexToAddress("0x0")

const ERC20ABI = `[
	{
		"constant":true,
		"inputs":[
				{
					"name":"tokenOwner",
					"type":"address"
				}
		],
		"name":"balanceOf",
		"outputs":[
				{
					"name":"balance",
					"type":"uint256"
				}
		],
		"payable":false,
		"stateMutability":"view",
		"type":"function"
	}
]`
