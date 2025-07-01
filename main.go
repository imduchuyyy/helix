package main

import (
	"fmt"

	"github.com/imduchuyyy/helix-wallet/actions"
	"github.com/imduchuyyy/helix-wallet/cli"
	"github.com/imduchuyyy/helix-wallet/keyring"
)

func main() {
	app := cli.NewCli()
	app.SetPrompt("Enter entropy to generate wallet > ")
	entropy, ok := app.AskEntropy()
	if !ok {
		return
	}
	keyring, err := keyring.New(entropy)
	if err != nil {
		fmt.Println("Error creating keyring:", err)
		return
	}
	address, err := keyring.GetEVMAddress()
	if err != nil {
		fmt.Println("Error generating EVM address:", err)
		return
	}
	fmt.Println("Login to Address:", address.Hex())

	action := actions.New(keyring)
	app.RegisterCommands(keyring.Commands())
	app.RegisterCommands(action.Commands())

	app.SetPrompt("Helix > ")

	app.Run()
}
