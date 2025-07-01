package main

import (
	"fmt"

	"github.com/imduchuyyy/crypto-lite/cli"
	"github.com/imduchuyyy/crypto-lite/keyring"
)

func main() {
	keyring := keyring.New()

	app := cli.NewCli()
	app.SetPrompt("Enter entropy to generate wallet > ")
	app.RegisterCommands(keyring.Commands())

	entropy, ok := app.AskEntropy()

	if !ok {
		return
	}

	fmt.Println(entropy)

	app.SetPrompt("Helix > ")

	app.Run()
}
