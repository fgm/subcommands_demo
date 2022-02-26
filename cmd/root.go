package cmd

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

type verboseKey struct{}

var VerboseKey = verboseKey{}

// Execute sets up the command chain and runs it.
func Execute(ctx context.Context) subcommands.ExitStatus {
	for _, command := range [...]struct {
		group string
		subcommands.Command
	}{
		{"help", subcommands.CommandsCommand()},  // Implement "commands"
		{"help", subcommands.FlagsCommand()},     // Implement "flags"
		{"help", subcommands.HelpCommand()},      // Implement "help"
		{"top", &top1{}},                         // Our first top-level command, without args
		{"top", &top2{}},                         // Our second top-level command, with args
		{"top", subcommands.Alias("1", &top1{})}, // An alias for our top1 command

	} {
		subcommands.Register(command.Command, command.group)
	}

	verbose := flag.Bool("v", false, "Be more verbose")
	flag.Parse()
	ctx = context.WithValue(ctx, VerboseKey, *verbose)

	return subcommands.Execute(ctx, "meaning")
}
