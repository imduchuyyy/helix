package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func main() {
	fmt.Println("SeedVault CLI - Interactive Mode")
	fmt.Println("Type 'help' to see available commands.")

	scanner := bufio.NewScanner(os.Stdin)
	var seed string

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}
		args := strings.Fields(input)
		switch args[0] {

		case "help":
			fmt.Println("Available commands:")
			fmt.Println("  init     - Generate a new BIP-39 mnemonic")
			fmt.Println("  backup   - Encrypt and backup current seed")
			fmt.Println("  split    - Split seed with Shamir's Secret Sharing")
			fmt.Println("  restore  - Restore seed from keystore/share files")
			fmt.Println("  decrypt  - Decrypt and display the seed (CAUTION)")
			fmt.Println("  exit     - Exit the CLI")

		case "init":
			// TODO: generate mnemonic and assign to `seed`
			fmt.Println("[init] Generating new mnemonic... (not yet implemented)")

		case "backup":
			if seed == "" {
				fmt.Println("No seed in memory. Please run 'init' or 'restore'.")
				break
			}
			pass := promptPassword("Enter passphrase to encrypt: ")
			confirm := promptPassword("Confirm passphrase: ")
			if pass != confirm {
				fmt.Println("Passphrase mismatch. Aborting.")
				break
			}
			// TODO: encrypt `seed` with KDF + AES-GCM and save keystore
			fmt.Println("[backup] Encrypt and backup not yet implemented.")

		case "split":
			// TODO: implement Shamir's Secret Sharing splitting `seed`
			fmt.Println("[split] Shamir share splitting not yet implemented.")

		case "restore":
			if len(args) < 2 {
				fmt.Println("Usage: restore <file1> [file2 ...]")
				break
			}
			pass := promptPassword("Enter passphrase to decrypt: ")
			// TODO: decrypt or reconstruct seed into `seed`
			fmt.Println("[restore] Restore not yet implemented.", pass)

		case "decrypt":
			pass := promptPassword("Enter passphrase to decrypt: ")
			// TODO: decrypt `seed` and print
			fmt.Println("[decrypt] Decrypt not yet implemented.", pass)

		case "exit", "quit":
			fmt.Println("Exiting SeedVault CLI. Goodbye!")
			return

		default:
			fmt.Printf("Unknown command: %s. Type 'help'.\n", args[0])
		}
	}
}

// promptPassword reads a hidden passphrase from stdin
func promptPassword(prompt string) string {
	fmt.Print(prompt)
	bytePass, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading passphrase:", err)
		os.Exit(1)
	}
	return string(bytePass)
}
