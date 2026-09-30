// Concept: cleaning input with TrimSpace and ReplaceAll
// Task: complete normalize so it trims surrounding whitespace and turns every tab into a single space
// Expected output: hello gostlings
// Hint: the strings package trims the ends of a string, and it replaces every
//       occurrence of one substring with another (Go doc: strings)
// Stuck?: strings.TrimSpace for the ends; strings.ReplaceAll for the tabs.

package main

import (
	"fmt"
	"strings"
)

func normalize(s string) string {
	// TODO: Trim surrounding whitespace, then replace tabs with spaces.
	return s
}

func main() {
	fmt.Println(normalize("  hello\tgostlings  "))
}
