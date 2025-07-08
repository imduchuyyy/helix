package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/imduchuyyy/helix-wallet/types"
)

// Cli represents the command-line interface
type Cli struct {
	scanner    *bufio.Scanner
	commands   map[string]types.Command
	prompt     string
	welcomeMsg string

	action types.Action
}

// NewCli creates a new CLI instance
func NewCli(action types.Action) *Cli {
	cli := &Cli{
		scanner:    bufio.NewScanner(os.Stdin),
		commands:   make(map[string]types.Command),
		prompt:     "> ",
		welcomeMsg: "Helix-wallet CLI - \nType 'help' to see available commands.",
		action:     action,
	}

	// Register built-in exit command
	cli.registerCommands(cli.Commands())
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

func (c *Cli) registerCommand(cmd types.Command) {
	if _, exists := c.commands[cmd.Name]; exists {
		fmt.Printf("Warning: Command '%s' is already registered and will be overwritten.\n", cmd.Name)
	}
	c.commands[cmd.Name] = cmd
}

// RegisterCommand adds a new command to the CLI
func (c *Cli) registerCommands(cmds []types.Command) {
	for _, cmd := range cmds {
		c.registerCommand(cmd)
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
			err := cmd.Handler(args[1:])
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Printf("Unknown command: %s. Type 'help' to see available commands.\n", cmdName)
		}
	}
}

// helpHandler displays help information for all commands
func (c *Cli) helpHandler(args []string) error {
	if len(args) > 0 {
		// Help for a specific command
		cmdName := args[0]
		if cmd, exists := c.commands[cmdName]; exists {
			if cmd.Usage != "" {
				fmt.Printf("%s: %s\nUsage: %s", cmd.Name, cmd.Description, cmd.Usage)
				return nil
			}
			fmt.Printf("%s: %s", cmd.Name, cmd.Description)
			return nil
		}
		return fmt.Errorf("unknown command: %s", cmdName)
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

	fmt.Println(helpText.String())

	return nil
}
