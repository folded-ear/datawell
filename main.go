/*
Package main is the main entrypoint to the datawell application, including
configuration loading.
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

var commands = []*Command{
	greetCmd,
	&serveCmd.Command,
	userListCmd,
	&userAddCmd.Command,
	&userEditCmd.Command,
	&userDeleteCmd.Command,
	demoCmd,
	echoCmd,
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 || args[0] == "-h" || args[0] == "-help" || args[0] == "--help" || args[0] == "help" {
		flag.Usage()
		return
	}

	var cmd *Command
	var prefixCmds []*Command
	name := args[0]
	for _, c := range commands {
		if c.Name == name {
			cmd = c
			break
		} else if strings.HasPrefix(c.Name, name) {
			prefixCmds = append(prefixCmds, c)
		}
	}

	if cmd == nil {
		if len(prefixCmds) == 1 {
			cmd = prefixCmds[0]
		} else {
			if len(prefixCmds) > 1 {
				names := []string{}
				for _, c := range prefixCmds {
					names = append(names, c.Name)
				}
				fmt.Printf("error: %q prefix-matches multiple commands: %v\n", name, names)
			} else {
				fmt.Printf("error: unknown command %q\n", name)
			}
			fmt.Printf("run 'datawell help' for usage\n")
			os.Exit(1)
		}
	}

	cmd.Exec(args[1:])
}

func usage() {
	fmt.Print(usagePrefix)
	flag.PrintDefaults()
	err := usageTmpl.Execute(os.Stdout, commands)
	if err != nil {
		fmt.Println(err)
	}
}

var usagePrefix = `
datawell is what EventLog became.

Usage:
    datawell [options] <subcommand> [subcommand options]

Options:
`

var usageTmpl = template.Must(template.New("usage").Parse(
	`
Subcommands:{{range .}}
{{.Name | printf "%-10s"}} {{.Summary}}
{{.SprintFlagDefaults}}{{end}}
`))
