// Concept: constants and iota
// Task: add the Green and Blue constants so they have the values 1 and 2
// Expected output: 0 1 2
// Hint: inside a const block, iota increments per line and an empty expression repeats the previous one (Go Tour: Basics 16)

package main

import "fmt"

const (
	Red = iota
	// TODO: Add the Green and Blue constants here, using iota.
)

func main() {
	fmt.Println(Red, Green, Blue)
}
