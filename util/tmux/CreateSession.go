package tmux

import (
	"os"
	"log"
	"os/exec"
)

func	 CreateSession(name string) {
	createCmd := exec.Command("tmux", "new-session", "-s", name)
	createCmd.Stdin = os.Stdin
	createCmd.Stdout = os.Stdout
	createCmd.Stderr = os.Stderr
	if err := createCmd.Run(); err != nil {
		log.Fatal(err)
	}
}
