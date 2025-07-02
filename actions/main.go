package actions

import (
	"fmt"

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
			Description: "All balances on [network]",
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

	tokenList, err := a.fetchTokenList(chain.TokenListURL)

	if err != nil {
		return "", fmt.Errorf("failed to fetch token list: %w", err)
	}

	fmt.Println("Fetched token list:", tokenList)

	return "Balance token", nil
}
