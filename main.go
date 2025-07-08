package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/imduchuyyy/helix-wallet/cli"
	"github.com/imduchuyyy/helix-wallet/evm"
	"github.com/imduchuyyy/helix-wallet/types"
	"golang.org/x/term"
)

var CHAIN = map[string]func(entropy string) (types.Action, error){
	"eth": func(entropy string) (types.Action, error) {
		return evm.New(entropy, "Ethereum", "https://eth.llamarpc.com", "https://raw.githubusercontent.com/Uniswap/default-token-list/refs/heads/main/src/tokens/mainnet.json")
	},
}

func askEntropy() (string, bool) {
	fmt.Print("Enter entropy: ")
	// Read password without echoing to terminal
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil || len(bytePassword) == 0 {
		fmt.Println("\nInvalid entropy input. Please try again.")
		return "", false
	}
	fmt.Println() // Add a newline after password input
	return string(bytePassword), true
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
	action, err := genActionFunc(entropy)
	if err != nil {
		fmt.Println("Error initializing action:", err)
		return
	}
	app := cli.NewCli(action)
	fmt.Println("Using chain:", action.ChainName())
	app.SetPrompt("Helix > ")

	app.Run()
}
