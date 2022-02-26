package cmd

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

type (
	debugKey   struct{}
	verboseKey struct{}
)

var (
	DebugKey   = debugKey{}
	VerboseKey = verboseKey{}
)

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

	debug := flag.Bool("debug", false, "Show debug information")
	verbose := flag.Bool("v", false, "Be more verbose")
	subcommands.ImportantFlag("v")
	flag.Parse()
	ctx = context.WithValue(ctx, DebugKey, *debug)
	ctx = context.WithValue(ctx, VerboseKey, *verbose)

	return subcommands.Execute(ctx, "meaning")
}
