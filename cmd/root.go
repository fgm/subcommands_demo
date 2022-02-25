package cmd

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

// Execute sets up the command chain and runs it.
func Execute(ctx context.Context) subcommands.ExitStatus {
	for _, command := range [...]subcommands.Command{
		subcommands.CommandsCommand(), // Implement "commands"
		subcommands.FlagsCommand(),    // Implement "flags"
		subcommands.HelpCommand(),     // Implement "help"
	} {
		subcommands.Register(command, "")
	}

	flag.Parse()
	return subcommands.Execute(ctx)
}
