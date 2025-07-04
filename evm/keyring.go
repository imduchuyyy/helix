package evm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func (e EVMAction) GetAddress() (string, error) {
	seed := crypto.Keccak256([]byte(e.entropy + "evm" + "helix-wallet"))

	privateKey, err := crypto.ToECDSA(seed)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return "", err
	}

	return crypto.PubkeyToAddress(privateKey.PublicKey).String(), nil
}
