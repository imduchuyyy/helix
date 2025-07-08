package evm

import (
	"fmt"
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/imduchuyyy/helix-wallet/common"
)

func (e EVMAction) GetAddress() (string, error) {
	seed := crypto.Keccak256([]byte(e.entropy + CHAIN_TYPE + common.WALLET_POSTFIX))

	privateKey, err := crypto.ToECDSA(seed)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return "", err
	}

	return crypto.PubkeyToAddress(privateKey.PublicKey).String(), nil
}

func (e EVMAction) GetPrivateKey() (string, error) {
	seed := crypto.Keccak256([]byte(e.entropy + CHAIN_TYPE + common.WALLET_POSTFIX))

	return ethcommon.BytesToHash(seed).String(), nil
}

func (e EVMAction) signTransaction(tx *ethtypes.Transaction, chainId *big.Int) (*ethtypes.Transaction, error) {
	seed := crypto.Keccak256([]byte(e.entropy + CHAIN_TYPE + common.WALLET_POSTFIX))
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
