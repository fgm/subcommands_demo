package cmd

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	"github.com/google/subcommands"
	"gopkg.in/yaml.v3"
)

type explainCommand visitCommand

func explainExplainCommand(commander *subcommands.Commander, w io.Writer) {
	commands := NewTop3(w, nil)
	fmt.Fprintln(w, "\nDemoes overriding ExplainCommand to describe top3.")
	fmt.Fprintln(w, "- Builtin version using private explain:")
	commander.ExplainCommand(w, commands)

	fmt.Fprintln(w, "- Custom version in YAML format:")
	commander.ExplainCommand = func(w io.Writer, c subcommands.Command) {
		fs := flag.NewFlagSet("commands", flag.ContinueOnError)
		c.SetFlags(fs)
		var flags []struct{ Name, Default, Usage string }
		fs.VisitAll(func(f *flag.Flag) {
			flags = append(flags, struct{ Name, Default, Usage string }{f.Name, f.DefValue, f.Usage})
		})
		_ = yaml.NewEncoder(w).Encode(map[string]any{
			c.Name(): map[string]any{
				"synopsis": c.Synopsis(),
				"usage":    c.Usage(),
				"flags":    flags,
			},
		})
	}
	commander.ExplainCommand(w, commands)
}

func explainExplainGroup(commander *subcommands.Commander, w io.Writer) {
	fmt.Fprintln(w, "\nDemoes overriding ExplainGroup.")
	fmt.Fprintln(w, "- Builtin version using private explainGroup:")
	commander.VisitGroups(func(group *subcommands.CommandGroup) {
		commander.ExplainGroup(w, group)
	})

	fmt.Fprintln(w, "- Custom version in YAML format, without access to group contents:")
	groups := map[string]int{}
	commander.ExplainGroup = func(w io.Writer, group *subcommands.CommandGroup) {
		groups[group.Name()] = group.Len()
	}
	commander.VisitGroups(func(group *subcommands.CommandGroup) {
		commander.ExplainGroup(w, group)
	})
	_ = yaml.NewEncoder(w).Encode(groups)
}

func explainExplain(commander *subcommands.Commander, w io.Writer) {
	fmt.Fprintln(w, "\nDemoes overriding Explain.")
	fmt.Fprintln(w, "- Builtin version using private commander.explain:")
	commander.Explain(w)

	fmt.Fprintln(w, "\n- Custom version, build from commander methods:")
	commander.Explain = func(w io.Writer) {
		fmt.Fprintln(w, "Use any commander.(Explain|Visit)* methods")
	}
	commander.Explain(w)
}

func (c explainCommand) Name() string                                   { return "explain" }
func (c explainCommand) Synopsis() string                               { return "demoes commander Explain* function fields" }
func (c explainCommand) Usage() string                                  { return c.Name() }
func (c *explainCommand) SetFlags(set *flag.FlagSet)                    {}
func (c *explainCommand) SetCommander(commander *subcommands.Commander) { c.commander = commander }
func (c *explainCommand) Execute(ctx context.Context, fs *flag.FlagSet, args ...any) subcommands.ExitStatus {
	commander, w := c.commander, c.outW
	explainExplainCommand(commander, w)
	explainExplainGroup(commander, w)
	explainExplain(commander, w)
	return subcommands.ExitSuccess
}

func NewExplain(outW io.Writer, logger *log.Logger) subcommands.Command {
	return &explainCommand{
		outW:   outW,
		logger: logger,
	}
}
