package keyring

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/imduchuyyy/helix-wallet/types"
)

type EVMKeyring struct {
	entropy string
}

func NewEVMKeyring(entropy string) (EVMKeyring, error) {
	return EVMKeyring{
		entropy: entropy,
	}, nil
}

func (k EVMKeyring) GetAddress() (string, error) {
	seed := crypto.Keccak256([]byte(k.entropy + "evm" + "helix-wallet"))

	privateKey, err := crypto.ToECDSA(seed)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return "", err
	}

	return crypto.PubkeyToAddress(privateKey.PublicKey).String(), nil
}

func (k *EVMKeyring) SignTransaction(tx *ethtypes.Transaction, chainId *big.Int) (*ethtypes.Transaction, error) {
	seed := crypto.Keccak256([]byte(k.entropy + "evm" + "helix-wallet"))
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

func (k *EVMKeyring) Commands() []types.Command {
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

func (k *EVMKeyring) handleGetAddress(args []string) (string, error) {
	address, err := k.GetAddress()
	if err != nil {
		return "", err
	}
	return address, nil
}

func (k *EVMKeyring) handleGetPrivateKey(args []string) (string, error) {
	seed := crypto.Keccak256([]byte(k.entropy + "evm" + "helix-wallet"))

	fmt.Println(common.Bytes2Hex(seed))

	return "", nil
}
