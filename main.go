package main

import (
	"fmt"

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
	fmt.Println(entropy)
	keyring := keyring.New(entropy)
	address := keyring.GetEVMAddress()
	fmt.Println("Generated EVM Address:", address.Hex())
	app.RegisterCommands(keyring.Commands())

	app.SetPrompt("Helix > ")

	app.Run()
}
