// Concept: building a string efficiently with strings.Builder
// Task: complete joinWords so it joins the words with a single space
// Expected output: hello gostlings
// Hint: repeated concatenation copies the whole string every time; the strings
//       package offers a builder that appends without that cost (Go doc: strings)
// Stuck?: strings.Builder's WriteString and WriteByte; append the separating
//       space yourself.

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
