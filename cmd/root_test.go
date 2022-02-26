package cmd

import (
	"context"
	"io"
	"log"
	"strings"
	"testing"

	"github.com/google/subcommands"
)

type describer func(io.Writer, *log.Logger) []description

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
		{"no actual args", nil, []string{arg0}, subcommands.ExitUsageError, "", ""},
		{"top1, undeclared global flag, no args", Describe, []string{arg0, "-bad", "top1"}, subcommands.ExitUsageError, "", "Error parsing CLI flags: flag provided but not defined: -bad\n"},
		{"top1, undeclared top1 flag, no args", Describe, []string{arg0, "top1", "-bad"}, subcommands.ExitUsageError, "", ""},
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
				t.Fatalf("Expected error:\n%s\ngot error:\n%s", check.expectedErr, actualErr)
			}
			if actualOut := outW.String(); actualOut != check.expectedOut {
				t.Fatalf("Expected output:\n%s\ngot output:\n%s", check.expectedOut, actualOut)
			}
		})
	}
}
