package cmd

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

// top1 is the type for a minimal command implementing subcommands.Command:
//   - no data stored
//   - no arguments
//   - no flags
type top1 struct{}

func (cmd *top1) Name() string {
	return "top1"
}

func (cmd *top1) Synopsis() string {
	return "top1 is an example top-level custom command without arguments"
}

func (cmd *top1) Usage() string {
	return cmd.Name()
}

func (cmd *top1) SetFlags(fs *flag.FlagSet) {}

func (cmd *top1) Execute(context.Context, *flag.FlagSet, ...interface{}) subcommands.ExitStatus {
	fmt.Printf("In %s\n", cmd.Name())
	return subcommands.ExitSuccess
}
