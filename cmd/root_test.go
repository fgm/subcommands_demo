package cmd

import (
	"context"
	"io"
	"log"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/subcommands"
)

type describer func(io.Writer, *log.Logger) []description

const (
	NoArgsErr = `Usage: Test_Execute <flags> <subcommand> <subcommand args>

Subcommands for help:
	commands         list all command names
	flags            describe all known top-level flags
	help             describe subcommands and their syntax


Top-level flags (use "Test_Execute flags" for a full list):
  -v=false: Be more verbose
`
	UndeclareGlobalFlagNoArgsErr = `flag provided but not defined: -bad
Usage: Test_Execute <flags> <subcommand> <subcommand args>

Subcommands for help:
	commands         list all command names
	flags            describe all known top-level flags
	help             describe subcommands and their syntax

Subcommands for top:
	top1, 1          top1 is an exemple top-level custom command without arguments
	top2             top2 is an exemple top-level custom command with arguments
	top3             top3 is an exemple top-level custom command with nested subcommands


Top-level flags (use "Test_Execute flags" for a full list):
  -v=false: Be more verbose
Error parsing CLI flags: flag provided but not defined: -bad
`
	UndeclaredTop1FlagNoArgsErr = `top1  -prefix string
    	Add a prefix to the result
`
)

// Test_Execute verifies the cmd.Execute function without exercising the commands themselves.
func Test_Execute(t *testing.T) {
	const arg0 = "Test_Execute"
	checks := [...]struct {
		name           string
		describe       describer
		args           []string
		expectedStatus subcommands.ExitStatus
		expectedOut    string
		expectedErr    string
	}{
		{"nil args", nil, nil, subcommands.ExitFailure, "", "Expected at least one argument for the program name, got none\n"},
		{"no actual args", nil, []string{arg0}, subcommands.ExitUsageError, "", NoArgsErr},
		{"top1, undeclared global flag, no args", Describe, []string{arg0, "-bad", "top1"}, subcommands.ExitUsageError, "", UndeclareGlobalFlagNoArgsErr},
		{"top1, undeclared top1 flag, no args", Describe, []string{arg0, "top1", "-bad"}, subcommands.ExitUsageError, "", UndeclaredTop1FlagNoArgsErr},
	}

	for _, check := range checks {
		check := check
		t.Run(check.name, func(t *testing.T) {
			t.Parallel()
			outW := &strings.Builder{}
			errW := &strings.Builder{}
			if actualStatus := Execute(context.Background(), outW, errW, check.args, 0, check.describe); actualStatus != check.expectedStatus {
				t.Fatalf("Expected: %d, got %d", check.expectedStatus, actualStatus)
			}
			if actualErr := errW.String(); actualErr != check.expectedErr {
				t.Fatal(cmp.Diff(check.expectedErr, actualErr))
			}
			if actualOut := outW.String(); actualOut != check.expectedOut {
				t.Fatalf("Expected output:\n%s\ngot output:\n%s", check.expectedOut, actualOut)
			}
		})
	}
}
