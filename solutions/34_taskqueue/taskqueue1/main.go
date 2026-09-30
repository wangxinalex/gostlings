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
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func run(args []string, out io.Writer, st Store) error {
	if len(args) == 0 {
		return errors.New("expected a subcommand: add, list, or done")
	}
	switch args[0] {
	case "add":
		flags := flag.NewFlagSet("add", flag.ContinueOnError)
		flags.SetOutput(io.Discard)
		priority := flags.Int("priority", 0, "task priority")
		if err := flags.Parse(args[1:]); err != nil {
			return err
		}
		if flags.NArg() != 1 {
			return errors.New("add takes exactly one title")
		}
		added, err := st.Add(flags.Arg(0), *priority)
		if err != nil {
			return err
		}
		fmt.Fprintln(out, added.Label())
		return nil
	case "list":
		for _, current := range st.List() {
			fmt.Fprintln(out, current.Label())
		}
		return nil
	case "done":
		if len(args) != 2 {
			return errors.New("done takes exactly one id")
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("invalid id %q: %w", args[1], err)
		}
		return st.Complete(id)
	default:
		return fmt.Errorf("unknown subcommand %q", args[0])
	}
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
