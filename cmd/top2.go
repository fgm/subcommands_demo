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

func top2Execute(ctx context.Context, cmd *top, fs *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	if ctx.Value(VerboseKey).(bool) {
		cmd.logger.Printf("In %s.\n", cmd.Name())
	}

	message := strings.Join(fs.Args(), " ")
	if cmd.prefix != "" {
		message = strings.Join(append([]string{cmd.prefix}, message), ": ")
	}
	fmt.Fprintln(cmd.outW, message)
	return subcommands.ExitSuccess
}

func NewTop2(outW io.Writer, logger *log.Logger) *top {
	const name = "top2"
	return &top{
		logger:   logger,
		name:     name,
		outW:     outW,
		prefix:   "",
		run:      top2Execute,
		synopsis: fmt.Sprintf("%s is an exemple top-level custom command with arguments", name),
		usage:    fmt.Sprintf("%s arg1 arg2 ...", name),
	}
}
