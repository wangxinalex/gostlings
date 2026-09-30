// Concept: checking substring membership with strings.Contains
// Task: complete contains so it reports whether text contains substr
// Expected output: true
// Hint: the strings package has a predicate that answers "does text contain
//       substr?" (Go doc: strings)
// Stuck?: strings.Contains; check its argument order and what it returns.

package main

import (
	"fmt"
	"strings"
)

func contains(text, substr string) bool {
	// TODO: Return whether text contains substr.
	return false
}

func main() {
	fmt.Println(contains("gostlings", "ling"))
}
