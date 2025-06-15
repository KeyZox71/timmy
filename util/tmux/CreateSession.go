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

func CreateDefaultSession(name string) {
	if _, err := os.Stat("flake.nix"); err == nil {
		createCmd := exec.Command("tmux", "new-session", "-s", name, "nix", "develop", "--command", "nvim", ";", "new-window")
		createCmd.Stdin = os.Stdin
		createCmd.Stdout = os.Stdout
		createCmd.Stderr = os.Stderr
		if err := createCmd.Run(); err != nil {
			log.Fatal(err)
		}
	} else if os.IsNotExist(err) {
		createCmd := exec.Command("tmux", "new-session", "-s", name, "nvim", ";", "new-window")
		createCmd.Stdin = os.Stdin
		createCmd.Stdout = os.Stdout
		createCmd.Stderr = os.Stderr
		if err := createCmd.Run(); err != nil {
			log.Fatal(err)
		}
	} else {
		// Handle unexpected errors
		log.Fatal(err)
	}
}
