package types

type Service interface {
	Commands() []Command
}
