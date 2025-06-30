package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/imduchuyyy/crypto-lite/types"
)

// CommandHandler represents a function that handles a specific command

// Cli represents the command-line interface
type Cli struct {
	scanner    *bufio.Scanner
	commands   map[string]types.Command
	prompt     string
	welcomeMsg string
}

// NewCli creates a new CLI instance
func NewCli() *Cli {
	cli := &Cli{
		scanner:    bufio.NewScanner(os.Stdin),
		commands:   make(map[string]types.Command),
		prompt:     "> ",
		welcomeMsg: "Crypto-Lite CLI - Interactive Mode\nType 'help' to see available commands.",
	}

	// Register built-in help command
	cli.RegisterCommand(types.Command{
		Name:        "help",
		Description: "Shows available commands",
		Handler:     cli.helpHandler,
	})

	// Register built-in exit command
	cli.RegisterCommand(types.Command{
		Name:        "exit",
		Description: "Exits the application",
		Handler: func(args []string) (string, error) {
			fmt.Println("Exiting. Goodbye!")
			os.Exit(0)
			return "", nil
		},
	})
	return cli
}

// SetPrompt changes the CLI prompt
func (c *Cli) SetPrompt(prompt string) {
	c.prompt = prompt
}

// SetWelcomeMessage changes the welcome message
func (c *Cli) SetWelcomeMessage(msg string) {
	c.welcomeMsg = msg
}

func (c *Cli) RegisterCommand(cmd types.Command) {
	if _, exists := c.commands[cmd.Name]; exists {
		fmt.Printf("Warning: Command '%s' is already registered and will be overwritten.\n", cmd.Name)
	}
	c.commands[cmd.Name] = cmd
}

// RegisterCommand adds a new command to the CLI
func (c *Cli) RegisterCommands(cmds []types.Command) {
	for _, cmd := range cmds {
		c.RegisterCommand(cmd)
	}
}

// Run starts the interactive CLI
func (c *Cli) Run() {
	fmt.Println(c.welcomeMsg)

	for {
		fmt.Print(c.prompt)
		if !c.scanner.Scan() {
			break
		}

		input := strings.TrimSpace(c.scanner.Text())
		if input == "" {
			continue
		}

		args := strings.Fields(input)
		cmdName := args[0]

		if cmd, exists := c.commands[cmdName]; exists {
			output, err := cmd.Handler(args[1:])
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else if output != "" {
				fmt.Println(output)
			}
		} else {
			fmt.Printf("Unknown command: %s. Type 'help' to see available commands.\n", cmdName)
		}
	}
}

// helpHandler displays help information for all commands
func (c *Cli) helpHandler(args []string) (string, error) {
	if len(args) > 0 {
		// Help for a specific command
		cmdName := args[0]
		if cmd, exists := c.commands[cmdName]; exists {
			if cmd.Usage != "" {
				return fmt.Sprintf("%s: %s\nUsage: %s", cmd.Name, cmd.Description, cmd.Usage), nil
			}
			return fmt.Sprintf("%s: %s", cmd.Name, cmd.Description), nil
		}
		return "", fmt.Errorf("unknown command: %s", cmdName)
	}

	// General help
	var helpText strings.Builder
	helpText.WriteString("Available commands:\n")

	// Get all commands and sort them alphabetically
	var cmdNames []string
	for name := range c.commands {
		cmdNames = append(cmdNames, name)
	}

	// We could sort cmdNames here for nicer output

	for _, name := range cmdNames {
		cmd := c.commands[name]
		helpText.WriteString(fmt.Sprintf("  %-12s - %s\n", name, cmd.Description))
	}

	return helpText.String(), nil
}
