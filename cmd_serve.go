package main

import (
	"fmt"
	"github.com/folded-ear/datawell/config"
)

var serveCmd = &Command{
	Name:    "serve",
	Usage:   "",
	Summary: "start the web server",
	Run:     serveRun,
}
var servePort uint

func init() {
	serveCmd.Flag.UintVar(&servePort, "port", 8080, "Port to listen on")
}

func serveRun(cmd *Command, args ...string) {
	config, _ := config.LoadConfig()
	fmt.Printf("env: %v\n", config.Env)
	fmt.Printf("dsn: %v\n", config.DataSourceName)
	fmt.Printf("run serve on port %v\n", servePort)
}
