package cmd

import (
	"context"
	"flag"
	"io"
	"log"

	"github.com/google/subcommands"
)

// top is the type for a reusable command implementing subcommands.Command:
type top struct {
	name, synopsis, usage string // Reuse support

	logger *log.Logger // Injected logger
	outW   io.Writer   // Injected writers

	prefix string // Actual features
	run    func(context.Context, *top, *flag.FlagSet, ...any) subcommands.ExitStatus
}

func (cmd top) Name() string {
	return cmd.name
}

func (cmd top) Synopsis() string {
	return cmd.synopsis
}

func (cmd top) Usage() string {
	return cmd.usage
}

func (cmd *top) SetFlags(fs *flag.FlagSet) {
	fs.StringVar(&cmd.prefix, "prefix", "", "Add a prefix to the result")
}

func (cmd *top) Execute(ctx context.Context, fs *flag.FlagSet, args ...any) subcommands.ExitStatus {
	if cmd.run == nil {
		// The logger was injected in Execute().
		cmd.logger.Printf("command %s is not runnable", cmd.name)
		return subcommands.ExitFailure
	}
	return cmd.run(ctx, cmd, fs, args)
}
