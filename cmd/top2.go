package cmd

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/google/subcommands"
)

func top2Execute(ctx context.Context, cmd *top, fs *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if ctx.Value(VerboseKey).(bool) {
		fmt.Printf("In %s.\n", cmd.Name())
	}
	fmt.Println(strings.Join(append([]string{cmd.prefix}, fs.Args()...), ": "))
	return subcommands.ExitSuccess
}

func NewTop2() *top {
	const name = "top2"
	return &top{
		name:     name,
		synopsis: fmt.Sprintf("%s is an exemple top-level custom command with arguments", name),
		usage:    fmt.Sprintf("%s arg1 arg2 ...", name),
		prefix:   "",
		run:      top2Execute,
	}
}
