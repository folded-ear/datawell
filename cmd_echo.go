package main

import "fmt"

var echoCmd = &Command{
	Name:    "echo",
	Usage:   "i echo back all flags and args",
	Summary: "echo out flags and args",
	Run:     echoRun,
}

func echoRun(cmd *Command, args ...string) {
	for i, s := range args {
		fmt.Printf("%2d) %v\n", i, s)
	}
}
