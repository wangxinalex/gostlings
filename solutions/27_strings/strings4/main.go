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
	var b strings.Builder
	for i, word := range words {
		if i > 0 {
			b.WriteByte(' ')
		}
		b.WriteString(word)
	}
	return b.String()
}

func main() {
	fmt.Println(joinWords([]string{"hello", "gostlings"}))
}
