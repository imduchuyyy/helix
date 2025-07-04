package main

import (
	"fmt"
	"os"

	"github.com/imduchuyyy/helix-wallet/cli"
	"github.com/imduchuyyy/helix-wallet/common"
	"github.com/imduchuyyy/helix-wallet/handler"
)

func main() {
	chainDenote := os.Getenv("CHAIN")
	app := cli.NewCli()
	app.SetPrompt("Enter entropy to generate wallet > ")
	entropy, ok := app.AskEntropy()
	if !ok {
		return
	}
	action, ok := common.GetChainAction(chainDenote, entropy)
	handler := handler.New(action)
	if !ok {
		fmt.Println("Invalid chain specified. Please set the CHAIN environment variable to a valid chain name.")
		return
	}
	fmt.Println("Using chain:", action.ChainName())
	app.SetPrompt("Helix > ")
	app.RegisterCommands(handler.Commands())

	app.Run()
}
