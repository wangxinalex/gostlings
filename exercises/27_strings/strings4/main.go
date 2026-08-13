// Concept: building a string efficiently with strings.Builder
// Task: complete joinWords so it joins the words with a single space
// Expected output: hello gostlings
// Hint: var b strings.Builder; use b.WriteString and b.WriteByte(' ') (Go doc: strings)

package main

import (
	"fmt"
	"strings"
)

func joinWords(words []string) string {
	// TODO: Build and return the space-joined string with strings.Builder.
	return ""
}

func main() {
	fmt.Println(joinWords([]string{"hello", "gostlings"}))
}
