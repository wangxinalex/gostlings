// Concept: multiple return values
// Task: complete the body of divide so it returns the quotient and the remainder
// Expected output: 3 1
// Hint: a function can return multiple results; the result types are listed in parentheses (Go Tour: Basics 6)

package main

import "fmt"

func divide(a, b int) (int, int) {
	// TODO: Return the quotient and the remainder of a divided by b.
}

func main() {
	q, r := divide(7, 2)
	fmt.Println(q, r)
}
