package types

type CommandHandler func(args []string) error

// Command represents a CLI command with its handler and description
type Command struct {
	Name        string
	Description string
	Handler     CommandHandler
	Usage       string
}
