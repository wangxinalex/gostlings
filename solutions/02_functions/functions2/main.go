// Concept: function parameters and arguments
// Task: fix this program so it compiles and runs
// Expected output: 3
// Hint: argument types must match the declared parameter types (Go Tour: Basics 5)

package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(1, 2))
}
