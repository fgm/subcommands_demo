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

// sub31Execute is called when top3 is invoked without arguments.
func sub31Execute(ctx context.Context, cmd *top, fs *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	if ctx.Value(VerboseKey).(bool) {
		cmd.logger.Printf("In %s.\n", cmd.Name())
	}
	message := fmt.Sprintf("hello %s", cmd.Name())
	if cmd.prefix != "" {
		message = strings.Join(append([]string{cmd.prefix}, message), ": ")
	}
	fmt.Fprintln(cmd.outW, message)
	return subcommands.ExitSuccess
}

func NewSub31(outW io.Writer, logger *log.Logger) subcommands.Command {
	const name = "sub31"
	return &top{
		logger:   logger,
		name:     name,
		outW:     outW,
		prefix:   "",
		run:      sub31Execute,
		synopsis: fmt.Sprintf("%s is an example subcommand", name),
		usage:    name,
	}
}
