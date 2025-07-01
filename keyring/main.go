package keyring

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/imduchuyyy/helix-wallet/types"
)

type Keyring struct {
	privateKey *ecdsa.PrivateKey
}

func New(entropy string) (*Keyring, error) {
	seed := crypto.Keccak256([]byte(entropy + "evm" + "helix-wallet"))

	privateKey, err := crypto.ToECDSA(seed)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return nil, err
	}

	return &Keyring{
		privateKey: privateKey,
	}, nil
}

func (k *Keyring) GetEVMAddress() (common.Address, error) {
	return crypto.PubkeyToAddress(k.privateKey.PublicKey), nil
}

func (k *Keyring) Commands() []types.Command {
	return []types.Command{
		{
			Name:        "address",
			Description: "Get Address",
			Handler:     k.handleGetAddress,
		},
	}
}
