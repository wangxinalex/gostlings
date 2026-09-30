// Concept: custom ordering with sort.Slice
// Task: complete byLength so it sorts words by length descending, ties alphabetically
// Expected output: [python go c]
// Hint: give the sort package a less function that decides the order of any two
//       positions; "less" means "comes first", so descending order inverts the
//       primary comparison and falls back to the other field for ties (Go doc: sort)
// Stuck?: sort.Slice takes the slice and a function of two indices returning bool.

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
