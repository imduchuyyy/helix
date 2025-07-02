package actions

import (
	"fmt"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/imduchuyyy/helix-wallet/common"
	"github.com/imduchuyyy/helix-wallet/keyring"
	"github.com/imduchuyyy/helix-wallet/types"
)

type Action struct {
	keyring *keyring.Keyring
}

func New(keyring *keyring.Keyring) *Action {
	return &Action{
		keyring: keyring,
	}
}

func (a *Action) Commands() []types.Command {
	return []types.Command{
		{
			Name:        "transfer",
			Description: "Transfer [amount] of [token] to [address] on [network]",
			Handler:     a.handleTransfer,
		},
		{
			Name:        "balance",
			Description: "All balances on [network]. Example: balance eth",
			Handler:     a.handleBalance,
		},
	}
}

func (a *Action) handleTransfer(args []string) (string, error) {
	return "Transfer token", nil
}

func (a *Action) handleBalance(args []string) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("network argument is required")
	}
	chain, ok := common.CHAIN[args[0]]
	if !ok {
		return "", fmt.Errorf("network %s not supported", args[0])
	}

	fmt.Printf("Fetching balances for network: %s\n", chain.Name)

	// address, err := a.keyring.GetEVMAddress()
	// if err != nil {
	// 	return "", fmt.Errorf("failed to get EVM address: %w", err)
	// }

	tokenWithBalance, err := a.fetchTokenBalances(chain.TokenListURL, chain.Rpcs[0], ethcommon.HexToAddress("0x4fff0f708c768a46050f9b96c46c265729d1a62f"))
	if err != nil {
		return "", fmt.Errorf("failed to fetch token list: %w", err)
	}

	for _, token := range tokenWithBalance {
		fmt.Printf("Token: %s, Balance: %s\n", token.Detail.Symbol, token.Balance)
	}

	return "", nil
}
