// Concept: the command layer wires flags, the store, and output together.
// Task: implement run so it dispatches the add, list, and done subcommands.
// Expected behavior: focused tests pass (run `go test ./exercises/34_taskqueue/taskqueue1`)
// Hint: give each subcommand its own flag set, and write user-facing output to the
//       writer you were handed instead of to the process's stdout, so a test can
//       read it. Validation belongs to the task type and lookup failures to the
//       store, so this layer reports what they return instead of duplicating the
//       checks; a returned error is what makes the process exit non-zero.
// Stuck?: flag.NewFlagSet with ContinueOnError keeps parsing errors as values;
//       strconv turns the id argument into a number.

package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func run(args []string, out io.Writer, st Store) error {
	// TODO: dispatch add, list, and done. add takes a priority flag and exactly one
	//       title, list prints the queue in order, and done takes exactly one id.
	return errors.New("run is not implemented")
}

func main() {
	queue := NewMemory()
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stdout, "usage: taskqueue <add|list|done> [flags]")
		return
	}
	if err := run(os.Args[1:], os.Stdout, queue); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
