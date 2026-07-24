// Concept: testing key presence with the comma-ok idiom
// Task: fix this program so it reports that "b" is missing from the map
// Expected output: b does not exist
// Hint: v, ok := m[key] sets ok to false when the key is absent (Go Tour: Moretypes 22)

package main

import "fmt"

func main() {
	m := map[string]int{"a": 1}
	// TODO: Use the two-value form of map lookup to check whether "b" exists.
	v := m["b"]
	fmt.Println("b exists with value", v)
}
