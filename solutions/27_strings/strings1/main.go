// Concept: checking substring membership with strings.Contains
// Task: complete contains so it reports whether text contains substr
// Expected output: true
// Hint: strings.Contains(text, substr) returns true when substr appears anywhere in text (Go doc: strings)

package main

import (
	"fmt"
	"strings"
)

func contains(text, substr string) bool {
	return strings.Contains(text, substr)
}

func main() {
	fmt.Println(contains("gostlings", "ling"))
}
