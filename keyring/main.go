package keyring

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/imduchuyyy/helix-wallet/types"
)

type Keyring struct {
	entropy string
	chain   types.Chain
}

func New(entropy string, chain types.Chain) (*Keyring, error) {
	return &Keyring{
		entropy: entropy,
		chain:   chain,
	}, nil
}

func (k *Keyring) GetAddress() (common.Address, error) {
	seed := crypto.Keccak256([]byte(k.entropy + "evm" + "helix-wallet"))

	privateKey, err := crypto.ToECDSA(seed)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return common.Address{}, err
	}

	return crypto.PubkeyToAddress(privateKey.PublicKey), nil
}

func (k *Keyring) Commands() []types.Command {
	return []types.Command{
		{
			Name:        "address",
			Description: "Get Address",
			Handler:     k.handleGetAddress,
		},
		{
			Name:        "private-key",
			Description: "Get Private Key",
			Handler:     k.handleGetPrivateKey,
		},
	}
}

func (k *Keyring) handleGetAddress(args []string) (string, error) {
	address, err := k.GetAddress()
	if err != nil {
		return "", err
	}
	return address.Hex(), nil
}

func (k *Keyring) handleGetPrivateKey(args []string) (string, error) {
	seed := crypto.Keccak256([]byte(k.entropy + "evm" + "helix-wallet"))

	fmt.Println(common.Bytes2Hex(seed))

	return "", nil
}
