// Concept: custom ordering with sort.Slice
// Task: complete byLength so it sorts words by length descending, ties alphabetically
// Expected output: [python go c]
// Hint: sort.Slice(words, func(i, j int) bool { ... }) (Go doc: sort)

package main

import (
	"fmt"
	"sort"
)

func byLength(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) != len(words[j]) {
			return len(words[i]) > len(words[j])
		}
		return words[i] < words[j]
	})
	return words
}

func main() {
	fmt.Println(byLength([]string{"c", "go", "python"}))
}
