package cmd

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"strings"
	"testing"

	"github.com/google/subcommands"
)

type expect struct {
	expectedSts subcommands.ExitStatus
	expectedErr string
	expectedOut string
}

// With this example implementation, there is no way this test should fail,
// but in more complex implementations with more features, it could be useful.
func TestTop_Properties(t *testing.T) {
	source := top{name: "a", synopsis: "b", usage: "c"}
	for _, check := range [...]struct {
		name   string
		field  string
		method func() string
	}{
		{"name", source.name, source.Name},
		{"synopsis", source.synopsis, source.Synopsis},
		{"usage", source.usage, source.Usage},
	} {
		if check.method() != check.field {
			t.Error(check.name)
		}
	}
}

func TestTop_SetFlags(t *testing.T) {
	const (
		testName = "TestTop_SetFlags"
		cmdName  = "a"
	)
	// This test verifies top.SetFlags, not top.Execute.
	runner := func(_ context.Context, cmd *top, _ *flag.FlagSet, _ ...any) subcommands.ExitStatus {
		fmt.Fprint(cmd.outW, cmd.prefix)
		return subcommands.ExitSuccess
	}
	checks := [...]struct {
		name string
		args []string
		expect
	}{
		{"no flags", []string{testName, cmdName}, expect{subcommands.ExitSuccess, "", ""}},
		{"bad flag", []string{testName, cmdName, "-bad"}, expect{subcommands.ExitUsageError, "", ""}},
		{"good flag", []string{testName, cmdName, "-prefix", "good"}, expect{subcommands.ExitSuccess, "", "good"}},
	}

	for _, check := range checks {
		check := check
		t.Run(check.name, func(t *testing.T) {
			t.Parallel()
			outW := &strings.Builder{}
			errW := &strings.Builder{}
			logger := log.New(errW, "", 0)
			fs := flag.NewFlagSet(testName, flag.ContinueOnError)
			cmd := top{name: cmdName, synopsis: "b", usage: "c", outW: outW, logger: logger, run: runner}
			commander := subcommands.NewCommander(fs, check.args[0])
			commander.Register(&cmd, "")

			if err := fs.Parse(check.args[1:]); err != nil {
				t.Errorf("parsing flags: %v", err)
			}
			if actualSts := commander.Execute(context.Background(), nil); actualSts != check.expectedSts {
				t.Fatalf("Expected %d, got %d", check.expectedSts, actualSts)
			}
			if actualErr := errW.String(); actualErr != check.expectedErr {
				t.Fatalf("Expected err:\n%sActual err:\n%s", check.expectedErr, actualErr)
			}
			if actualOut := outW.String(); actualOut != check.expectedOut {
				t.Fatalf("Expected out:\n%sActual out:\n%s", check.expectedOut, actualOut)
			}
		})
	}
}

func TestTop_Execute(t *testing.T) {
	const (
		cmdName  = "a"
		testName = "TestTop_Execute"
	)

	// This test verifies top.Execute, not top.SetFlags.
	NewNonRunnable := func(writer io.Writer, logger *log.Logger) *top {
		return &top{logger: logger, name: cmdName, outW: writer, synopsis: "b", usage: "b"}
	}

	checks := [...]struct {
		name       string
		args       []string
		cmdFactory func(io.Writer, *log.Logger) *top
		expect
	}{
		{"non runnable", []string{testName, cmdName}, NewNonRunnable, expect{subcommands.ExitFailure, fmt.Sprintf("command %s is not runnable\n", cmdName), ""}},
		{"top1 without args", []string{testName, "top1"}, NewTop1, expect{subcommands.ExitSuccess, "", "hello\n"}},
		{"top1 with args", []string{testName, "top1", "bad"}, NewTop1, expect{subcommands.ExitFailure, "top1 expects no arguments, called with 1: [bad]\n", ""}},
	}

	for _, check := range checks {
		check := check
		t.Run(check.name, func(t *testing.T) {
			t.Parallel()
			outW := &strings.Builder{}
			errW := &strings.Builder{}
			logger := log.New(errW, "", 0)
			cmd := check.cmdFactory(outW, logger)
			fs := flag.NewFlagSet(cmd.Name(), flag.PanicOnError)
			fs.Parse(check.args[2:])

			ctx := context.WithValue(context.Background(), VerboseKey, false)

			if actualSts := cmd.Execute(ctx, fs, nil); actualSts != check.expectedSts {
				t.Fatalf("Expected %d, got %d", check.expectedSts, actualSts)
			}
			if actualErr := errW.String(); actualErr != check.expectedErr {
				t.Fatalf("Expected err:\n%sActual err:\n%s", check.expectedErr, actualErr)
			}
			if actualOut := outW.String(); actualOut != check.expectedOut {
				t.Fatalf("Expected out:\n%sActual out:\n%s", check.expectedOut, actualOut)
			}
		})
	}
}

type topCheck struct {
	name    string
	verbose bool
	prefix  string
	args    []string
	expect
}

func baseTopNTest(t *testing.T, check topCheck, factory func(writer io.Writer, logger *log.Logger) *top) {
	t.Run(check.name, func(t *testing.T) {
		t.Parallel()
		outW := &strings.Builder{}
		errW := &strings.Builder{}
		logger := log.New(errW, "", 0)
		cmd := factory(outW, logger)
		ctx := context.WithValue(context.Background(), VerboseKey, check.verbose)

		fs := flag.NewFlagSet(cmd.Name(), flag.PanicOnError)
		cmd.SetFlags(fs) // Normally performed by commander.Execute()
		args := check.args
		if check.prefix != "" {
			args = append([]string{"-prefix", check.prefix}, args...)
		}
		fs.Parse(args)

		if actualSts := cmd.Execute(ctx, fs, nil); actualSts != check.expectedSts {
			t.Fatalf("Expected %d, got %d", check.expectedSts, actualSts)
		}
		if actualErr := errW.String(); actualErr != check.expectedErr {
			t.Fatalf("Expected err:\n%sActual err:\n%s", check.expectedErr, actualErr)
		}
		if actualOut := outW.String(); actualOut != check.expectedOut {
			t.Fatalf("Expected out:\n%sActual out:\n%s", check.expectedOut, actualOut)
		}
	})
}
