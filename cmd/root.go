package cmd

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

// Execute sets up the command chain and runs it.
func Execute(ctx context.Context) subcommands.ExitStatus {
	for _, command := range [...]subcommands.Command{
		subcommands.CommandsCommand(),   // Implement "commands"
		subcommands.FlagsCommand(),      // Implement "flags"
		subcommands.HelpCommand(),       // Implement "help"
		&top1{},                         // Our first top-level command, without args
		&top2{},                         // Our second top-level command, with args
		subcommands.Alias("1", &top1{}), // An alias for our top1 command
	} {
		subcommands.Register(command, "")
	}

	flag.Parse()
	return subcommands.Execute(ctx, "meaning", 42)
}
