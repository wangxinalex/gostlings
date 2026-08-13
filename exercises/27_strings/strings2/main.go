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
	// TODO: Return parts joined with sep.
	return ""
}

func splitParts(s, sep string) []string {
	// TODO: Return s split on sep.
	return nil
}

func main() {
	fmt.Println(joinParts([]string{"go", "rust", "python"}, ","))
	fmt.Println(strings.Join(splitParts("go rust python", " "), " "))
}
