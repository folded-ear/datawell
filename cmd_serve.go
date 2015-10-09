package main

import (
	"fmt"
	"github.com/folded-ear/datawell/config"
)

type ServeCommand struct {
	Command
	servePort *uint
}

var serveCmd = &ServeCommand{
	Command: Command{
		Name:    "serve",
		Usage:   "",
		Summary: "start the web server",
	},
}

func init() {
	serveCmd.Run = serveRun
	serveCmd.servePort = serveCmd.Flag.Uint("port", 8080, "Port to listen on")
}

func serveRun(cmd *Command, args ...string) {
	config := config.LoadConfig()
	fmt.Printf("env: %v\n", config.Env)
	fmt.Printf("dsn: %v\n", config.DataSourceName)
	fmt.Printf("run serve on port %v\n", *serveCmd.servePort)
}
