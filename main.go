package main

import (
	"fmt"
	"os"

	"github.com/imduchuyyy/helix-wallet/actions"
	"github.com/imduchuyyy/helix-wallet/cli"
	"github.com/imduchuyyy/helix-wallet/common"
	"github.com/imduchuyyy/helix-wallet/keyring"
)

func main() {
	chainDenote := os.Getenv("CHAIN")

	chain, ok := common.CHAIN[chainDenote]
	if !ok {
		fmt.Println("Invalid chain specified. Please set the CHAIN environment variable to a valid chain name.")
		return
	}
	fmt.Println("Using chain:", chain.Name)

	app := cli.NewCli()
	app.SetPrompt("Enter entropy to generate wallet > ")
	entropy, ok := app.AskEntropy()
	if !ok {
		return
	}
	keyring, err := keyring.New(entropy, chain)
	if err != nil {
		fmt.Println("Error creating keyring:", err)
		return
	}
	address, err := keyring.GetAddress()
	if err != nil {
		fmt.Println("Error generating wallet:", err)
		return
	}
	fmt.Println("Login to Address:", address.Hex())

	action := actions.New(keyring, chain)
	app.RegisterCommands(keyring.Commands())
	app.RegisterCommands(action.Commands())

	app.SetPrompt("Helix > ")

	app.Run()
}
