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
	// TODO: Sort words by length descending, ties alphabetically.
	return words
}

func main() {
	fmt.Println(byLength([]string{"c", "go", "python"}))
}
