package cmd

import (
	"context"
	"flag"
	"io"
	"log"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/subcommands"
)

func subTest(t *testing.T,
	factory func(io.Writer, *log.Logger) subcommands.Command,
	verbose bool,
	prefix string,
	run func(context.Context, *top, *flag.FlagSet, ...any) subcommands.ExitStatus,
	expect expect,
) {
	ctx := context.WithValue(context.Background(), VerboseKey, verbose)
	outW := &strings.Builder{}
	errW := &strings.Builder{}
	logger := log.New(errW, "", 0)
	cmd := factory(outW, logger).(*top)
	cmd.prefix = prefix
	run(ctx, cmd, nil, nil) // These functions always return ExitSuccess.
	if actualErr := errW.String(); actualErr != expect.expectedErr {
		t.Fatalf(cmp.Diff(actualErr, expect.expectedErr))
	}
	if actualOut := outW.String(); actualOut != expect.expectedOut {
		t.Fatalf(cmp.Diff(actualOut, expect.expectedOut))
	}
}

func Test_sub31Execute(t *testing.T) {
	checks := [...]struct {
		name    string
		verbose bool
		prefix  string
		expect
	}{
		{"quiet, no prefix", false, "", expect{expectedErr: "", expectedOut: "hello sub31\n"}},
		{"verbose, no prefix", true, "", expect{expectedErr: "In sub31.\n", expectedOut: "hello sub31\n"}},
		{"quiet, prefix", false, "foo", expect{expectedErr: "", expectedOut: "foo: hello sub31\n"}},
		{"verbose, prefix", true, "foo", expect{expectedErr: "In sub31.\n", expectedOut: "foo: hello sub31\n"}},
	}
	for _, check := range checks {
		check := check
		t.Run(check.name, func(t *testing.T) {
			t.Parallel()
			subTest(t, NewSub31, check.verbose, check.prefix, sub31Execute, check.expect)
		})
	}
}
