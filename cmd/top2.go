package cmd

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/google/subcommands"
)

// top2 is the type for a simple command implementing subcommands.Command::
//   - no data stored
//   - takes arguments
//   - one flag
type top2 struct {
	prefix string
}

func (cmd *top2) Name() string {
	return "top2"
}

func (cmd *top2) Synopsis() string {
	return "top2 is an example top-level custom command with arguments"
}

func (cmd *top2) Usage() string {
	return fmt.Sprintf("%s arg1 arg2 ...", cmd.Name())
}

func (cmd *top2) SetFlags(fs *flag.FlagSet) {
	fs.StringVar(&cmd.prefix, "prefix", "", "Add a prefix to the result")
}

func (cmd *top2) Execute(ctx context.Context, fs *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if ctx.Value(VerboseKey).(bool) {
		fmt.Printf("In %s.\n", cmd.Name())
	}
	fmt.Println(strings.Join(append([]string{cmd.prefix}, fs.Args()...), ": "))
	return subcommands.ExitSuccess
}
