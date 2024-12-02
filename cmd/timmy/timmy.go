package main

import (
	"os"

	"github.com/keyzox71/timmy/cmd/ts"
)

func	main() {
	args := os.Args

	if len(args) <= 1 {
		helpMsg()
		return
	}

	if args[1] == "help" {
		helpMsg()
	} else if args[1] == "search" {
		ts.Ts()
	}
}
