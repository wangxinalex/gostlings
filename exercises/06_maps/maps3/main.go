// Concept: deleting map entries and ranging over a map
// Task: delete "a" from the map, then print each remaining entry as "k v"
// Expected output: b 2
// Hint: delete(m, key) removes an entry; for k, v := range m visits the rest (Go Tour: Moretypes 22; range: Moretypes 16)

package main

import "fmt"

func main() {
	m := map[string]int{"a": 1, "b": 2}
	// TODO: Delete "a" from m, then range over m printing each key and value on its own line.
}
