// Concept: handling a non-zero exit with exec.ExitError
// Task: complete exitCode so it returns the command's exit code, or an error if it failed to start
// Expected behavior: exitCode returns (0, nil) on success and (3, nil) for a command that exits 3
// Hint: err := cmd.Run(); if exitErr, ok := err.(*exec.ExitError); ok { return exitErr.ExitCode(), nil } (Go doc: os/exec)

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func exitCode(cmd *exec.Cmd) (int, error) {
	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return exitErr.ExitCode(), nil
		}
		return 0, err
	}
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
