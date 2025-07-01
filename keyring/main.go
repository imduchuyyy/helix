package keyring

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/imduchuyyy/helix-wallet/types"
)

type Keyring struct {
	entropy string
}

func New(entropy string) *Keyring {
	return &Keyring{
		entropy: entropy,
	}
}

func (w *Keyring) GetEVMAddress() common.Address {
	seed := crypto.Keccak256([]byte(w.entropy + "evm" + "helix-wallet"))
	fmt.Println("Seed:", common.BytesToHash(seed).String())
	return common.Address{}
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
