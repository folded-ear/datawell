/*
Package main is the main entrypoint to the datawell application, including
configuration loading.
*/
package main

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

const (
	DefaultDriverName     = "use db/dbconf.yml"
	DefaultDataSourceName = "use db/dbconf.yml"
)

var (
	env            = flag.String("env", "development", "the DB/Goose environment to use")
	driverName     = flag.String("db-driver-name", DefaultDriverName, "The database driver name (the first param to sql.Open)")
	dataSourceName = flag.String("db-datasource-name", DefaultDataSourceName, "The datasource name (the second param to sql.Open)")
)

var commands = []*Command{
	greetCmd,
	serveCmd,
	demoCmd,
}

type config struct {
	env            string
	driverName     string
	dataSourceName string
}

func newConfigFromFlagsAndGoose() (*config, error) {
	config := config{
		env: *env,
	}

	config.driverName = *driverName
	config.dataSourceName = *dataSourceName

	if *driverName == DefaultDriverName || *dataSourceName == DefaultDataSourceName {
		gDbConf, err := goose.NewDBConf("./db", *env, "")
		if err != nil {
			return nil, err
		}
		if *driverName == DefaultDriverName {
			config.driverName = gDbConf.Driver.Name
		}
		if *dataSourceName == DefaultDataSourceName {
			config.dataSourceName = gDbConf.Driver.OpenStr
		}
	}

	return &config, nil
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
{{.PrintFlagDefaultsToString}}{{end}}
`))
