package cmd

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

// top2 is the type for a simple command implementing subcommands.Command::
//   - no data stored
//   - takes arguments
//   - no flags
type top2 struct{}

func (cmd *top2) Name() string {
	return "top2"
}

func (cmd *top2) Synopsis() string {
	return "top2 is an example top-level custom command with arguments"
}

func (cmd *top2) Usage() string {
	return fmt.Sprintf("%s arg1 arg2 ...", cmd.Name())
}

func (cmd *top2) SetFlags(fs *flag.FlagSet) {}

func (cmd *top2) Execute(_ context.Context, fs *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	// Notice how the command line arguments are taken on the flag set, not on the variadic.
	fmt.Printf("In %s %v\n", cmd.Name(), fs.Args())
	return subcommands.ExitSuccess
}