// Concept: command-line flags with flag.NewFlagSet
// Task: complete parseArgs so it reads -name, -count, and -verbose
// Expected behavior: parseArgs returns the flag values and no error for valid input
// Hint: fs := flag.NewFlagSet("app", flag.ContinueOnError); register String/Int/Bool flags, then fs.Parse(args) (Go doc: flag)

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
