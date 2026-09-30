// Concept: running an external command with exec.Command and Output
// Task: complete echo so it runs "echo" with args and returns the trimmed output
// Expected output: hello gostlings
// Hint: the os/exec package builds a command from a program name and its
//       arguments, and a helper collects its standard output (Go doc: os/exec)
// Stuck?: exec.Command and its Output method; strings.TrimSpace drops the
//       trailing newline.

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func echo(args ...string) (string, error) {
	// TODO: Run echo with args and return the trimmed output.
	return "", nil
}

func main() {
	out, err := echo("hello", "gostlings")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(out)
}
