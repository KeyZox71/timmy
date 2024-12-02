package util

import (
	"os"
	"log"
	"path/filepath"
	"strings"
)

func	GetCurrentPathN(n int) string {
	pwd, err := os.Getwd()
	if (err != nil) { log.Fatal(err) }
	parts := strings.Split(filepath.Clean(pwd), string(os.PathSeparator))
	if (n > len(parts)) {
		n = len(parts)
	}
	path := strings.Join(parts[len(parts)-n:], "/")
	return path
}

