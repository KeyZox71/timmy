package tm

import (
	"os"

	path "github.com/keyzox71/timmy/util/path"
	tmux "github.com/keyzox71/timmy/util/tmux"
)

func	Tm() {
	if (len(os.Args) <= 2) {
		tmux.CreateSession(path.GetCurrentPathN(1))
		return
	}
	args := os.Args[2:]
	if (args[0] == "-2") {
		tmux.CreateSession(path.GetCurrentPathN(2))
		return
	} else {
		tmux.CreateSession(args[0])
	}
}
