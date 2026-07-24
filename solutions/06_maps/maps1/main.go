// Concept: adding entries to a map
// Task: add the key "apple" with value 3 to the map
// Expected output: 3
// Hint: insert or update an entry with m[key] = value (Go Tour: Moretypes 22)

package main

import "fmt"

func main() {
	m := map[string]int{}
	m["apple"] = 3
	fmt.Println(m["apple"])
}
