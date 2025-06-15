package main

import (
	"fmt"
)

func	helpMsg() {
	fmt.Println("timmy's help page\n")
	fmt.Println("USAGE: timmy <command> <args>\n")
	fmt.Println("COMMAND:")
	fmt.Println("  search - used to fuzzy find session")
	fmt.Println("  create - used to create session")
	fmt.Println("         - '-2' used to create a session named after the last two folder of the working directory")
	fmt.Println("         - 'name' used to create a session named with the name specified")
	fmt.Println("         - '-d' used to create a \"default workspace\" (with nvim as first tab and a shell as second)")
	fmt.Println("  version - used to print version number")
}

func	versionMsg() {
	fmt.Println("timmy version v0.1.1")
}
