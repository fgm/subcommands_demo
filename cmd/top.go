package cmd

import (
	"context"
	"flag"
	"log"

	"github.com/google/subcommands"
)

// top is the type for a reusable command implementing subcommands.Command:
type top struct {
	name, synopsis, usage string // Reuse support
	prefix                string // Actual features
	run                   func(context.Context, *top, *flag.FlagSet, ...interface{}) subcommands.ExitStatus
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

func (cmd *top) Execute(ctx context.Context, fs *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if cmd.run == nil {
		log.Printf("command %s is not runnable", cmd.name)
		return subcommands.ExitFailure
	}
	return cmd.run(ctx, cmd, fs, args)
}
