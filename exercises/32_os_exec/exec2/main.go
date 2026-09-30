// Concept: handling a non-zero exit with exec.ExitError
// Task: complete exitCode so it returns the command's exit code, or an error if it failed to start
// Expected behavior: exitCode returns (0, nil) on success and (3, nil) for a command that exits 3
// Hint: a command that started and exited non-zero reports a typed error that
//       carries the code; any other error means it never ran (Go doc: os/exec)
// Stuck?: exec.ExitError and its ExitCode method, matched with errors.As.

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func exitCode(cmd *exec.Cmd) (int, error) {
	// TODO: Run cmd and return its exit code, or the error when it does not start.
	return 0, nil
}

func main() {
	code, err := exitCode(exec.Command("sh", "-c", "exit 3"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("exit", code)
}
