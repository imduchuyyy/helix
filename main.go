package main

import (
	"github.com/imduchuyyy/crypto-lite/cli"
	"github.com/imduchuyyy/crypto-lite/keyring"
)

func main() {
	keyring := keyring.New()

	app := cli.NewCli()
	app.SetPrompt("crypto-lite > ")

	app.RegisterCommands(keyring.Commands())

	app.Run()
}
