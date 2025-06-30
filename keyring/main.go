package keyring

import "github.com/imduchuyyy/crypto-lite/types"

type Keyring struct {
}

func New() *Keyring {
	return &Keyring{}
}

func (w *Keyring) Commands() []types.Command {
	return []types.Command{
		{
			Name:        "create",
			Description: "Create a new wallet",
			Handler:     w.handleCreateWallet,
		},
	}
}
