//go:build !go1.18

package cmd

// any is only defined in Go 1.18+, so alias it manually for earlier versions.
type any = interface{}
