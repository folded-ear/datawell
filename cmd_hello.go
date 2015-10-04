package main

import "fmt"

var greetCmd = &Command{
	Name:    "greet",
	Usage:   "i say hello",
	Summary: "say hello",
	Run:     greetRun,
}

func greetRun(cmd *Command, args ...string) {
	var name string
	if len(args) == 0 {
		name = "world"
	} else {
		name = args[0]
	}
	fmt.Printf("Hello, %s!\n", name)
}
