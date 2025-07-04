package evm

import (
	"fmt"
	"math/big"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
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

func (e EVMAction) signTransaction(tx *ethtypes.Transaction, chainId *big.Int) (*ethtypes.Transaction, error) {
	seed := crypto.Keccak256([]byte(e.entropy + "evm" + "helix-wallet"))
	privateKey, err := crypto.ToECDSA(seed)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return nil, err
	}

	signedTx, err := ethtypes.SignTx(tx, ethtypes.NewEIP155Signer(chainId), privateKey)

	if err != nil {
		fmt.Println("Error signing transaction:", err)
		return nil, err
	}
	return signedTx, nil
}
