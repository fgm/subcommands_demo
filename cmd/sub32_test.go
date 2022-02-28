package cmd

import (
	"testing"
)

func Test_sub32Execute(t *testing.T) {
	checks := [...]struct {
		name    string
		verbose bool
		prefix  string
		expect
	}{
		{"quiet, no prefix", false, "", expect{expectedErr: "", expectedOut: "hello sub32\n"}},
		{"verbose, no prefix", true, "", expect{expectedErr: "In sub32.\n", expectedOut: "hello sub32\n"}},
		{"quiet, prefix", false, "foo", expect{expectedErr: "", expectedOut: "foo: hello sub32\n"}},
		{"verbose, prefix", true, "foo", expect{expectedErr: "In sub32.\n", expectedOut: "foo: hello sub32\n"}},
	}
	for _, check := range checks {
		check := check
		t.Run(check.name, func(t *testing.T) {
			t.Parallel()
			subTest(t, NewSub32, check.verbose, check.prefix, sub32Execute, check.expect)
		})
	}
}
