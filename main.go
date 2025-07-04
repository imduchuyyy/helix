package main

import (
	"fmt"
	"os"

	"github.com/imduchuyyy/helix-wallet/cli"
	"github.com/imduchuyyy/helix-wallet/evm"
	"github.com/imduchuyyy/helix-wallet/types"
)

var CHAIN = map[string]func(entropy string) types.Action{
	"eth": func(entropy string) types.Action {
		return evm.New(entropy, "Ethereum", "https://eth.llamarpc.com", "https://raw.githubusercontent.com/Uniswap/default-token-list/refs/heads/main/src/tokens/mainnet.json")
	},
}

func askEntropy() (string, bool) {
	var entropy string
	fmt.Print("Enter entropy: ")
	_, err := fmt.Scanln(&entropy)
	if err != nil || entropy == "" {
		fmt.Println("Invalid entropy input. Please try again.")
		return "", false
	}
	return entropy, true
}

func main() {
	chainDenote := os.Getenv("CHAIN")
	genActionFunc, exists := CHAIN[chainDenote]
	if !exists {
		fmt.Println("Invalid or missing CHAIN environment variable. Please set it to a valid chain name (e.g., 'eth').")
		return
	}
	entropy, ok := askEntropy()
	if !ok {
		return
	}
	action := genActionFunc(entropy)
	app := cli.NewCli(action)
	fmt.Println("Using chain:", action.ChainName())
	app.SetPrompt("Helix > ")

	app.Run()
}
