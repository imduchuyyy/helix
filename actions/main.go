package actions

import (
	"fmt"
	"math"
	"math/big"

	"github.com/imduchuyyy/helix-wallet/keyring"
	"github.com/imduchuyyy/helix-wallet/types"
)

type Action struct {
	keyring *keyring.Keyring
	chain   types.Chain
}

func New(keyring *keyring.Keyring, chain types.Chain) *Action {
	return &Action{
		keyring: keyring,
		chain:   chain,
	}
}

func (a *Action) Commands() []types.Command {
	return []types.Command{
		{
			Name:        "transfer",
			Description: "Transfer [amount] of [token] to [address]",
			Handler:     a.handleTransfer,
		},
		{
			Name:        "balance",
			Description: "All balances",
			Handler:     a.handleBalance,
		},
	}
}

func (a *Action) handleTransfer(args []string) (string, error) {
	return "Transfer token", nil
}

func (a *Action) handleBalance(args []string) (string, error) {
	address, err := a.keyring.GetEVMAddress()
	if err != nil {
		return "", fmt.Errorf("failed to get EVM address: %w", err)
	}

	tokenWithBalance, err := a.fetchTokenBalances(a.chain.TokenListURL, a.chain.Rpcs[0], address)
	if err != nil {
		return "", fmt.Errorf("failed to fetch token list: %w", err)
	}

	for _, token := range tokenWithBalance {
		decimalBalance := new(big.Float).Quo(
			new(big.Float).SetInt(token.Balance),
			new(big.Float).SetFloat64(math.Pow10(int(token.Detail.Decimals))),
		)
		fmt.Printf("Token: %s, Balance: %.6f\n", token.Detail.Symbol, decimalBalance)
	}

	return "", nil
}
