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

func top1Execute(ctx context.Context, cmd *top, fs *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	if ctx.Value(VerboseKey).(bool) {
		cmd.logger.Printf("In %s.\n", cmd.Name())
	}
	if l := fs.NArg(); l != 0 {
		cmd.logger.Printf("%s expects no arguments, called with %d: %v", cmd.Name(), l, fs.Args())
		return subcommands.ExitFailure
	}
	message := "hello"
	if cmd.prefix != "" {
		message = strings.Join(append([]string{cmd.prefix}, message), ": ")
	}
	fmt.Fprintln(cmd.outW, message)
	return subcommands.ExitSuccess
}

func NewTop1(outW io.Writer, logger *log.Logger) *top {
	const name = "top1"
	return &top{
		logger:   logger,
		name:     name,
		outW:     outW,
		prefix:   "",
		run:      top1Execute,
		synopsis: fmt.Sprintf("%s is an exemple top-level custom command without arguments", name),
		usage:    name,
	}
}
