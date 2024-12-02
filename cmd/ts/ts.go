package ts

import (
	"log"
	"os"
	"os/exec"
	"strings"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
)

func	Ts() {
	cmd, err := exec.Command("tmux", "list-sessions", "-F", "#{session_name}").Output()
	if err != nil {
		log.Fatal(err)
	}
	str := strings.Split(string(cmd), "\n")
	res, err := fuzzyfinder.FindMulti(
		str,
		func(i int) string{
			return str[i]
		},
	)
	if (err != nil) {
		log.Print(err)
		return
	}
	attachCmd := exec.Command("tmux", "attach", "-t", string(str[res[0]]))
	attachCmd.Stderr = os.Stderr
	attachCmd.Stdout = os.Stdout
	attachCmd.Stdin = os.Stdin

	if err = attachCmd.Run(); err != nil {
		log.Fatal(err)
	}
}
