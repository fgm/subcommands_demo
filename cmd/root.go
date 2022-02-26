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
	commander := subcommands.DefaultCommander

	for _, command := range [...]struct {
		group string
		subcommands.Command
	}{
		{"help", commander.CommandsCommand()},      // Implement "commands"
		{"help", commander.FlagsCommand()},         // Implement "flags"
		{"help", commander.HelpCommand()},          // Implement "help"
		{"top", NewTop1()},                         // Our first top-level command, without args
		{"top", NewTop2()},                         // Our second top-level command, with args
		{"top", subcommands.Alias("1", NewTop1())}, // An alias for our top1 command

	} {
		commander.Register(command.Command, command.group)
	}

	debug := flag.Bool("debug", false, "Show debug information")
	verbose := flag.Bool("v", false, "Be more verbose")
	commander.ImportantFlag("v")
	flag.Parse()

	ctx = context.WithValue(ctx, DebugKey, *debug)
	ctx = context.WithValue(ctx, VerboseKey, *verbose)

	return commander.Execute(ctx, "meaning")
}
