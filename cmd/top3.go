package cmd

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/google/subcommands"
)

func top3Internal(_ context.Context, cmd *top, _ *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	message := "hello top3"
	if cmd.prefix != "" {
		message = strings.Join(append([]string{cmd.prefix}, message), ": ")
	}
	_, _ = fmt.Fprintln(cmd.outW, message)
	return subcommands.ExitSuccess
}

func top3Execute(ctx context.Context, cmd *top, topFS *flag.FlagSet, args ...any) subcommands.ExitStatus {
	name := cmd.Name()
	if ctx.Value(VerboseKey).(bool) {
		cmd.logger.Printf("In %s.\n", cmd.Name())
	}
	// Handle command called without subcommands.
	if topFS.NArg() == 0 {
		return top3Internal(ctx, cmd, topFS, args)
	}

	// Create a flag.FlagSet from args to use only remaining args
	// Continue on error to support testing.
	fs := flag.NewFlagSet(cmd.Name(), flag.ContinueOnError)
	fs.SetOutput(cmd.logger.Writer())

	// Create a custom commander to restart evaluation below this command.
	commander := subcommands.NewCommander(fs, name)
	commander.Output = cmd.outW
	commander.Error = cmd.logger.Writer()

	descriptions := []description{
		{name, commander.CommandsCommand()}, // Implement "commands"
		{name, commander.FlagsCommand()},    // Implement "flags"
		{name, commander.HelpCommand()},     // Implement "help"
		{name, NewSub31(cmd.outW, cmd.logger)},
		{name, NewSub32(cmd.outW, cmd.logger)},
	}
	for _, command := range descriptions {
		commander.Register(command.command, command.group)
	}

	// At this point, flags have been consumed, so fs.Parse() cannot fail.
	_ = fs.Parse(topFS.Args())

	return commander.Execute(ctx, fs)
}

func NewTop3(outW io.Writer, logger *log.Logger) *top {
	const name = "top3"
	return &top{
		logger:   logger,
		name:     name,
		outW:     outW,
		prefix:   "",
		run:      top3Execute,
		synopsis: fmt.Sprintf("%s is an exemple top-level custom command with nested subcommands", name),
		usage:    name,
	}
}
