package cmd

import (
	"testing"

	"github.com/google/subcommands"
)

func Test_top1Execute(t *testing.T) {
	const baseOut = "hello\n"
	checks := [...]topCheck{
		{"no args, no prefix, quiet", false, "", nil, expect{subcommands.ExitSuccess, "", baseOut}},
		{"no args, no prefix, verbose", true, "", nil, expect{subcommands.ExitSuccess, "In top1.\n", baseOut}},
		{"no args, prefix", false, "prefix", nil, expect{subcommands.ExitSuccess, "", "prefix: " + baseOut}},
		{"args", false, "", []string{"bad"}, expect{subcommands.ExitFailure, "top1 expects no arguments, called with 1: [bad]\n", ""}},
	}

	for _, check := range checks {
		baseTopNTest(t, check, NewTop1)
	}

}
