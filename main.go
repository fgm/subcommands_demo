package main

import (
	"context"
	"log"
	"os"

	"github.com/fgm/subcommands_demo/cmd"
)

func main() {
	sts := cmd.Execute(context.Background(), os.Stdout, os.Stderr, os.Args, log.LstdFlags, cmd.Describe)
	os.Exit(int(sts))
}
