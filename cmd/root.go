package cmd

import (
	"context"
	"flag"
	"io"
	"log"

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

type description struct {
	group   string
	command subcommands.Command
}

// Describe provides the list of non-builtin command descriptions,
// so that Execute can receive different sets of commands during tests.
func Describe(outW io.Writer, logger *log.Logger) []description {
	return []description{
		{"top", NewTop1(outW, logger)},                         // Our first top-level command, without args
		{"top", NewTop2(outW, logger)},                         // Our second top-level command, with args
		{"top", subcommands.Alias("1", NewTop1(outW, logger))}, // An alias for our top1 command
		{"top", NewTop3(outW, logger)},                         // Our command with subcommands
	}
}

// Execute sets up the command chain and runs it.
// It does not depend on any global nor mutates any.
func Execute(ctx context.Context,
	outW io.Writer, // Standard output for command results
	errW io.Writer, // Error output for logs
	args []string, // CLI args to avoid depending on the flag global
	logFlags int, // Log flags to make error message testing easier
	describe func(outW io.Writer, logger *log.Logger) []description, // Command registration descriptions
) subcommands.ExitStatus {
	// Do not depend on log.Default().
	logger := log.New(errW, "", logFlags)

	if len(args) < 1 {
		logger.Printf("Expected at least one argument for the program name, got none")
		return subcommands.ExitFailure
	}

	// Create a flag.FlagSet from args to avoid depending on global os.Args.
	// Continue on error to support testing instead of the ExitOnError on flag.CommandLine
	fs := flag.NewFlagSet(args[0], flag.ContinueOnError)

	// Create a custom commander to avoid depending on global flag.CommandLine and os.Args
	commander := subcommands.NewCommander(fs, args[0])

	descriptions := []description{
		{"help", commander.CommandsCommand()}, // Implement "commands"
		{"help", commander.FlagsCommand()},    // Implement "flags"
		{"help", commander.HelpCommand()},     // Implement "help"
	}
	if describe != nil {
		descriptions = append(descriptions, describe(outW, logger)...)
	}
	for _, command := range descriptions {
		commander.Register(command.command, command.group)
	}

	debug := fs.Bool("debug", false, "Show debug information")
	verbose := fs.Bool("v", false, "Be more verbose")
	commander.ImportantFlag("v")

	// Parse must not receive the program name, hence the slice.
	if err := fs.Parse(args[1:]); err != nil {
		logger.Printf("Error parsing CLI flags: %v", err)
		return subcommands.ExitUsageError
	}

	ctx = context.WithValue(ctx, DebugKey, *debug)
	ctx = context.WithValue(ctx, VerboseKey, *verbose)

	return commander.Execute(ctx, "meaning", 42)
}
