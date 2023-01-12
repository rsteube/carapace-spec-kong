package spec

import (
	"fmt"

	"github.com/alecthomas/kong"
	"gopkg.in/yaml.v2"
)

func Scrape(ctx kong.Context) {
	cmd := command(*ctx.Model.Node)
	m, err := yaml.Marshal(cmd)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(m))
}

func command(node kong.Node) Command {
	cmd := Command{
		Name:        node.Name,
		Aliases:     node.Aliases,
		Description: node.Help,
		Flags:       make(map[string]string),
		Commands:    make([]Command, 0),
	}

	if group := node.Group; group != nil {
		cmd.Group = group.Key
	}

	for _, flag := range node.Flags {
		formatted := ""

		if flag.Short != 0 {
			formatted += fmt.Sprintf("-%v, ", string(flag.Short))
		}
		formatted += fmt.Sprintf("--%v", flag.Name)

		switch {
		case flag.IsBool():
		//case optionalArgument:
		//	formatted += "?"
		default:
			formatted += "="
		}

		if flag.IsCounter() || flag.IsCumulative() {
			formatted += "*"
		}
		cmd.Flags[formatted] = flag.Help

		// TODO enum
	}

	for _, subcmd := range node.Children {
		if !subcmd.Hidden {
			cmd.Commands = append(cmd.Commands, command(*subcmd))
		}
	}
	return cmd
}