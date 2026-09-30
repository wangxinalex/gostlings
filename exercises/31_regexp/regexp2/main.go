// Concept: compiling a regexp and finding all matches
// Task: complete findNumbers so it returns every digit sequence in s
// Expected output: [123 45]
// Hint: compile the pattern first, because that step can fail, then ask the
//       compiled pattern for every non-overlapping match (Go doc: regexp)
// Stuck?: regexp.Compile returns the compiled pattern and an error; the find-all
//       method takes a limit where -1 means "all".

package main

import (
	"fmt"
	"regexp"
)

func findNumbers(s string) []string {
	// TODO: Return every digit sequence in s.
	return nil
}

func main() {
	fmt.Println(findNumbers("a=123 b=45"))
}
