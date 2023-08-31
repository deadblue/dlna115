package app

type Command interface {
	// Command name
	Name() string
	// Command description
	Desc() string
	// Init command with args
	Init(args []string) error
	// Print usage of command
	PrintUsage(prog string)
	// Run this command
	Run() error
}
