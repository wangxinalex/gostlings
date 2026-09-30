// Concept: command-line flags with flag.NewFlagSet
// Task: complete parseArgs so it reads -name, -count, and -verbose
// Expected behavior: parseArgs returns the flag values and no error for valid input
// Hint: build a private flag set that reports errors instead of exiting, register
//       one flag per type, then parse the argument slice (Go doc: flag)
// Stuck?: flag.NewFlagSet with flag.ContinueOnError, plus the String, Int, and
//       Bool registration methods.

package main

import (
	"flag"
	"fmt"
	"os"
)

func parseArgs(args []string) (name string, count int, verbose bool, err error) {
	fs := flag.NewFlagSet("app", flag.ContinueOnError)
	// TODO: Register the -name, -count, and -verbose flags, then parse args
	//       and return their values.
	_ = fs
	return "", 0, false, nil
}

func main() {
	name, count, verbose, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	fmt.Printf("name=%s count=%d verbose=%t\n", name, count, verbose)
}
