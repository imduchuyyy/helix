package onchain

import (
	"github.com/imduchuyyy/helix-wallet/keyring"
	"github.com/imduchuyyy/helix-wallet/types"
)

type OnchainAction struct {
	keyring *keyring.Keyring
}

func New(keyring *keyring.Keyring) *OnchainAction {
	return &OnchainAction{
		keyring: keyring,
	}
}

func (o *OnchainAction) Commands() []types.Command {
	return []types.Command{
		{
			Name:        "transfer",
			Description: "Get Address",
			Handler:     o.handleTransfer,
		},
	}
}

func (o *OnchainAction) handleTransfer(args []string) (string, error) {
	return "Transfer token", nil
}
