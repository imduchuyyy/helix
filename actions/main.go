package actions

import (
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
	}
}

func (a *Action) handleTransfer(args []string) (string, error) {
	return "Transfer token", nil
}
