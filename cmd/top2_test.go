package cmd

import (
	"testing"

	"github.com/google/subcommands"
)

func Test_top2Execute(t *testing.T) {
	const baseOut = "\n"
	checks := [...]topCheck{
		{"no args, no prefix, quiet", false, "", []string{}, expect{subcommands.ExitSuccess, "", baseOut}},
		{"no args, no prefix, verbose", true, "", []string{}, expect{subcommands.ExitSuccess, "In top2.\n", baseOut}},
		{"no args, prefix", false, "prefix", []string{}, expect{subcommands.ExitSuccess, "", "prefix: " + baseOut}},
		{"args", false, "", []string{"meaning", "42"}, expect{subcommands.ExitSuccess, "", "meaning 42\n"}},
	}

	for _, check := range checks {
		baseTopNTest(t, check, NewTop2)
	}

}
