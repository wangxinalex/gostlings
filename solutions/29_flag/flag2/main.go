// Concept: flags mixed with positional arguments
// Task: complete parseCommand so it reads -port and returns the leftover positional args in order
// Expected behavior: parseCommand returns the port and the positional args
// Hint: after fs.Parse, fs.Args() returns the non-flag arguments (Go doc: flag)

package main

import (
	"flag"
	"fmt"
	"os"
)

func parseCommand(args []string) (port int, rest []string, err error) {
	fs := flag.NewFlagSet("cmd", flag.ContinueOnError)
	portPtr := fs.Int("port", 0, "port")
	if err := fs.Parse(args); err != nil {
		return 0, nil, err
	}
	return *portPtr, fs.Args(), nil
}

func main() {
	port, rest, err := parseCommand(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	fmt.Printf("port=%d rest=%v\n", port, rest)
}
