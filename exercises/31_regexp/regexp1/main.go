// Concept: checking a pattern with regexp.MatchString
// Task: complete isHex so it reports whether s is a two-digit hex code
// Expected output: true false
// Hint: match the whole string rather than a substring of it, so both ends of the
//       pattern are anchored (Go doc: regexp)
// Stuck?: regexp.MatchString returns a bool and an error; the class covers 0-9
//       and a-f, twice.

package main

import (
	"fmt"
	"regexp"
)

func isHex(s string) bool {
	// TODO: Return whether s is exactly two hex digits.
	return false
}

func main() {
	fmt.Println(isHex("3f"), isHex("xyz"))
}
