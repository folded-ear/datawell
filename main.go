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
	serveCmd,
	demoCmd,
	echoCmd,
	listUsersCmd,
	addUserCmd,
	editUserCmd,
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
	name := args[0]
	for _, c := range commands {
		if strings.HasPrefix(c.Name, name) {
			cmd = c
			break
		}
	}

	if cmd == nil {
		fmt.Printf("error: unknown command %q\n", name)
		flag.Usage()
		os.Exit(1)
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
Commands:{{range .}}
{{.Name | printf "%-10s"}} {{.Summary}}
{{.SprintFlagDefaults}}{{end}}
`))
