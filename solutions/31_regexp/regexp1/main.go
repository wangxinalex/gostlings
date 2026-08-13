// Concept: checking a pattern with regexp.MatchString
// Task: complete isHex so it reports whether s is a two-digit hex code
// Expected output: true false
// Hint: regexp.MatchString(`^[0-9a-fA-F]{2}$`, s) (Go doc: regexp)

package main

import (
	"fmt"
	"regexp"
)

func isHex(s string) bool {
	matched, err := regexp.MatchString(`^[0-9a-fA-F]{2}$`, s)
	return err == nil && matched
}

func main() {
	fmt.Println(isHex("3f"), isHex("xyz"))
}
