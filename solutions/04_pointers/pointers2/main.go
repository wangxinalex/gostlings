// Concept: passing pointers to functions
// Task: implement setZero so the program prints 0
// Expected output: 0
// Hint: assigning through *p changes the caller's variable, not a copy (Go Tour: Methods 5)

package main

import "fmt"

func setZero(p *int) {
	*p = 0
}

func main() {
	n := 5
	setZero(&n)
	fmt.Println(n)
}
