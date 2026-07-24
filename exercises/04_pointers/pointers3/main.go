// Concept: allocating memory with new
// Task: fix this program so it prints 0 instead of panicking with a nil pointer dereference
// Expected output: 0
// Hint: a nil *int points to nothing; new(int) allocates a zeroed int and returns its pointer (new has no dedicated tour page; closest: Go Tour: Moretypes 1)

package main

import "fmt"

func main() {
	var p *int
	// TODO: Use new to make p point to a real int, so the next line prints 0 instead of panicking.
	fmt.Println(*p)
}
