package main

import "fmt"

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
	config, _ := newConfigFromFlagsAndGoose()
	fmt.Printf("env: %v\n", config.env)
	fmt.Printf("dsn: %v\n", config.dataSourceName)
	fmt.Printf("run serve on port %v\n", servePort)
}
