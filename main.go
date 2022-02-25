package main

import (
	"context"
	"os"

	"github.com/fgm/subcommands_demo/cmd"
)

func main() {
	ctx := context.Background()
	sts := cmd.Execute(ctx)
	os.Exit(int(sts))
}
