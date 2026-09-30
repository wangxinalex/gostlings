// Concept: flags mixed with positional arguments
// Task: complete parseCommand so it reads -port and returns the leftover positional args in order
// Expected behavior: parseCommand returns the port and the positional args
// Hint: after parsing, the flag set still holds the arguments that were not
//       flags (Go doc: flag)
// Stuck?: the flag set's Args method returns the positionals in order.

package main

import (
	"flag"
	"fmt"
	"os"
)

func parseCommand(args []string) (port int, rest []string, err error) {
	fs := flag.NewFlagSet("cmd", flag.ContinueOnError)
	// TODO: Register -port, parse args, then return the port and fs.Args().
	_ = fs
	return 0, nil, nil
}

func main() {
	port, rest, err := parseCommand(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	fmt.Printf("port=%d rest=%v\n", port, rest)
}
