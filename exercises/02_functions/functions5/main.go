// Concept: named return values
// Task: complete countdown so it assigns 3 to its named return value and uses a bare return
// Expected output: 3
// Hint: `(n int)` names the result; assign to n and use `return` without an expression (Go Tour: Basics 7)

package main

import "fmt"

func countdown() (n int) {
	// TODO: Assign 3 to n and return it with a bare return.
	return
}

func main() {
	fmt.Println(countdown())
}
