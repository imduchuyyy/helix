package common

import (
	"github.com/ethereum/go-ethereum/common"
)

var ZERO_ADDRESS = common.HexToAddress("0x0")

const WALLET_POSTFIX = "helix-wallet"

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
