package cli

import (
	"fmt"
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

func (c *Cli) Commands() []types.Command {
	return []types.Command{
		{
			Name:        "get_address",
			Description: "Get the wallet address",
			Handler:     c.handleGetAddress,
			Usage:       "get_address",
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
