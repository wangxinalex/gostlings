// Concept: joining and splitting strings
// Task: complete joinParts and splitParts
// Expected output: go,rust,python
// go rust python
// Hint: strings.Join(parts, sep) and strings.Split(s, sep) (Go doc: strings)

package main

import (
	"fmt"
	"strings"
)

func joinParts(parts []string, sep string) string {
	return strings.Join(parts, sep)
}

func splitParts(s, sep string) []string {
	return strings.Split(s, sep)
}

func main() {
	fmt.Println(joinParts([]string{"go", "rust", "python"}, ","))
	fmt.Println(strings.Join(splitParts("go rust python", " "), " "))
}
