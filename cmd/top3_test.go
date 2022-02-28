package cmd

import (
	"testing"

	"github.com/google/subcommands"
)

func Test_top3Execute(t *testing.T) {
	const (
		baseOut             = "hello top3\n"
		ArgNotSubcommandErr = `Usage: top3 <flags> <subcommand> <subcommand args>

Subcommands for top3:
	commands         list all command names
	flags            describe all known top-level flags
	help             describe subcommands and their syntax
	sub31            sub31 is an example subcommand
	sub32            sub32 is an example subcommand

`
	)
	checks := [...]topCheck{
		{"no args, no prefix, quiet", false, "", nil, expect{subcommands.ExitSuccess, "", baseOut}},
		{"no args, no prefix, verbose", true, "", nil, expect{subcommands.ExitSuccess, "In top3.\n", baseOut}},
		{"no args, prefix", false, "prefix", nil, expect{subcommands.ExitSuccess, "", "prefix: " + baseOut}},
		{"arg not subcommand", false, "", []string{"bad"}, expect{subcommands.ExitUsageError, ArgNotSubcommandErr, ""}},
	}

	for _, check := range checks {
		check := check
		t.Run(check.name, func(t *testing.T) {
			t.Parallel()
			baseTopNTest(t, check, NewTop3)
		})
	}

}
