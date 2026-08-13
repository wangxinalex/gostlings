// Concept: cleaning input with TrimSpace and ReplaceAll
// Task: complete normalize so it trims surrounding whitespace and turns every tab into a single space
// Expected output: hello gostlings
// Hint: strings.TrimSpace(s) and strings.ReplaceAll(s, old, new) (Go doc: strings)

package main

import (
	"fmt"
	"strings"
)

func normalize(s string) string {
	return strings.ReplaceAll(strings.TrimSpace(s), "\t", " ")
}

func main() {
	fmt.Println(normalize("  hello\tgostlings  "))
}
