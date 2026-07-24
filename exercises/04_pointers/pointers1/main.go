// Concept: declaring pointers and dereferencing them
// Task: fix this program so it prints 42, the value the pointer points to
// Expected output: 42
// Hint: *p reads the value a pointer p points to; p alone is just a memory address (Go Tour: Moretypes 1)

package main

import "fmt"

func main() {
	n := 42
	p := &n
	// TODO: Print the value p points to (42) instead of the pointer itself.
	fmt.Println(p)
}
