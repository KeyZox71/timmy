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
}
