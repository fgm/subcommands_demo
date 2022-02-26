package cmd

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/google/subcommands"
)

func top1Execute(ctx context.Context, cmd *top, fs *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if ctx.Value(VerboseKey).(bool) {
		fmt.Printf("In %s.\n", cmd.Name())
	}
	fmt.Println(strings.Join(append([]string{cmd.prefix}, "hello"), ": "))
	return subcommands.ExitSuccess
}

func NewTop1() *top {
	const name = "top1"
	return &top{
		name:     name,
		synopsis: fmt.Sprintf("%s is an exemple top-level custom command without arguments", name),
		usage:    name,
		prefix:   "",
		run:      top1Execute,
	}
}
