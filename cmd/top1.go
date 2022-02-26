package cmd

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/google/subcommands"
)

// top1 is the type for a minimal command implementing subcommands.Command:
//   - no data stored
//   - no arguments
//   - one flag
type top1 struct {
	prefix string
}

func (cmd *top1) Name() string {
	return "top1"
}

func (cmd *top1) Synopsis() string {
	return "top1 is an example top-level custom command without arguments"
}

func (cmd *top1) Usage() string {
	return cmd.Name()
}

func (cmd *top1) SetFlags(fs *flag.FlagSet) {
	fs.StringVar(&cmd.prefix, "prefix", "", "Add a prefix to the result")
}

func (cmd *top1) Execute(ctx context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if ctx.Value(VerboseKey).(bool) {
		fmt.Printf("In %s.\n", cmd.Name())
	}
	fmt.Println(strings.Join(append([]string{cmd.prefix}, "hello"), ": "))
	return subcommands.ExitSuccess
}
