// Concept: running an external command with exec.Command and Output
// Task: complete echo so it runs "echo" with args and returns the trimmed output
// Expected output: hello gostlings
// Hint: out, err := exec.Command("echo", args...).Output(); strings.TrimSpace(string(out)) (Go doc: os/exec)

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
