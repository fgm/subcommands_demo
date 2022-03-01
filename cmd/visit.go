package cmd

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"strings"
	"text/tabwriter"

	"github.com/google/subcommands"
)

type visitCommand struct {
	commander *subcommands.Commander
	logger    *log.Logger // Injected logger
	outW      io.Writer   // Injected writers
}

func visitAll(commander *subcommands.Commander, w io.Writer) {
	fmt.Fprintln(w, "VisitAll show all the commander flags:")
	tw := tabwriter.NewWriter(w, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(tw, "\tName\tDefault\tValue\tUsage\t")
	commander.VisitAll(func(f *flag.Flag) {
		fmt.Fprintf(tw, "\t%s\t%s\t%s\t%s\t\n", f.Name, f.DefValue, f.Value, f.Usage)
	})
	tw.Flush()
	fmt.Fprintln(w)
}

func visitAllImportant(commander *subcommands.Commander, w io.Writer) {
	fmt.Fprintln(w, "VisitAllImportant only shows the \"important\" flags:")
	tw := tabwriter.NewWriter(w, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(tw, "\tName\tDefault\tValue\tUsage\t")
	commander.VisitAllImportant(func(f *flag.Flag) {
		fmt.Fprintf(tw, "\t%s\t%s\t%s\t%s\t\n", f.Name, f.DefValue, f.Value, f.Usage)
	})
	tw.Flush()
	fmt.Fprintln(w)
}

func visitGroups(commander *subcommands.Commander, w io.Writer) {
	fmt.Fprintln(w, "VisitGroups only visits the command groups, not the commands:")
	tw := tabwriter.NewWriter(w, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(tw, "\tName\tLen\t")
	commander.VisitGroups(func(group *subcommands.CommandGroup) {
		fmt.Fprintf(tw, "\t%s\t%d\t\n", group.Name(), group.Len())
	})
	tw.Flush()
	fmt.Fprintln(w)
}

func visitCommands(commander *subcommands.Commander, w io.Writer) {
	fmt.Fprintln(w, "VisitCommands visits the commands themselves:")
	tw := tabwriter.NewWriter(w, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(tw, "\tGroup\tName\tSynopsis\tFlags\t")
	commander.VisitCommands(func(group *subcommands.CommandGroup, c subcommands.Command) {
		fs := flag.NewFlagSet("visit", flag.ContinueOnError)
		c.SetFlags(fs)
		var flags []string
		fs.VisitAll(func(f *flag.Flag) { flags = append(flags, f.Name) })
		fmt.Fprintf(tw, "\t%s\t%s\t%s\t%s\t\n", group.Name(), c.Name(), c.Synopsis(), strings.Join(flags, ", "))
	})
	tw.Flush()
	fmt.Fprintln(w, "(command usage omitted for readability)")
}

func (c visitCommand) Name() string                                   { return "visit" }
func (c visitCommand) Synopsis() string                               { return "demoes commander Visit* functions" }
func (c visitCommand) Usage() string                                  { return c.Name() }
func (c *visitCommand) SetFlags(set *flag.FlagSet)                    {}
func (c *visitCommand) SetCommander(commander *subcommands.Commander) { c.commander = commander }
func (c *visitCommand) Execute(ctx context.Context, fs *flag.FlagSet, args ...any) subcommands.ExitStatus {
	commander, w := c.commander, c.outW
	visitAll(commander, w)
	visitAllImportant(commander, w)
	visitGroups(commander, w)
	visitCommands(commander, w)
	return subcommands.ExitSuccess
}

func NewVisit(outW io.Writer, logger *log.Logger) subcommands.Command {
	return &visitCommand{
		outW:   outW,
		logger: logger,
	}
}
