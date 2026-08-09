// Concept: named return values
// Task: complete countdown so it assigns 3 to its named return value and uses a bare return
// Expected output: 3
// Hint: `(n int)` names the result; assign to n and use `return` without an expression (Go Tour: Basics 7)

package main

import "fmt"

func countdown() (n int) {
	n = 3
	return
}

func main() {
	fmt.Println(countdown())
}
