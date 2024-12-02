package tm

import (
	"log"
	"os"
	"os/exec"

	timmyPath "github.com/keyzox71/timmy/util/path"
)

func	 createSession(name string) {
	createCmd := exec.Command("tmux", "new-session", "-s", name)
	createCmd.Stdin = os.Stdin
	createCmd.Stdout = os.Stdout
	createCmd.Stderr = os.Stderr
	if err := createCmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func	Tm() {
	if (len(os.Args) <= 2) {
		createSession(timmyPath.GetCurrentPathN(1))
		return
	}
	args := os.Args[2:]
	if (args[0] == "-2") {
		createSession(timmyPath.GetCurrentPathN(2))
		return
	}
}
