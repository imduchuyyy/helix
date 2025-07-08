package cli

import (
	"fmt"
	"math"
	"math/big"
	"os"

	"github.com/imduchuyyy/helix-wallet/types"
)

func (c *Cli) handleGetAddress(args []string) error {
	address, err := c.action.GetAddress()
	if err != nil {
		return err
	}
	println("Wallet address:", address)
	return nil
}

func (c *Cli) handleGetPrivateKey(args []string) error {
	privateKey, err := c.action.GetPrivateKey()
	if err != nil {
		return err
	}
	println("Wallet private key:", privateKey)
	return nil
}

func (c *Cli) handleGetBalance(args []string) error {
	balances, err := c.action.GetTokenBalances()
	if err != nil {
		return err
	}

	if len(balances) == 0 {
		println("No tokens found in the wallet.")
		return nil
	}

	for _, token := range balances {
		decimalBalance := new(big.Float).Quo(
			new(big.Float).SetInt(token.Balance),
			new(big.Float).SetFloat64(math.Pow10(int(token.Decimals))),
		)
		fmt.Printf("Token: %s (%s), Balance: %s\n", token.Name, token.Symbol, decimalBalance.String())
	}
	return nil
}

func (c *Cli) Commands() []types.Command {
	return []types.Command{
		{
			Name:        "address",
			Description: "Get the wallet address",
			Handler:     c.handleGetAddress,
		},
		{
			Name:        "privatekey",
			Description: "Get the wallet private key",
			Handler:     c.handleGetPrivateKey,
		},
		{
			Name:        "balance",
			Description: "Get the wallet balance",
			Handler:     c.handleGetBalance,
		},
		{
			Name:        "help",
			Description: "Shows available commands",
			Handler:     c.helpHandler,
		},
		{
			Name:        "exit",
			Description: "Exits the application",
			Handler: func(args []string) error {
				fmt.Println("Exiting. Goodbye!")
				os.Exit(0)
				return nil
			},
		},
	}
}
