// Concept: variadic functions
// Task: complete the body of sum so it adds up all its arguments
// Expected output: 6
// Hint: variadic parameters collect arguments into a slice; use range to iterate over it (builds on Go Tour: Basics 4-7; variadic parameters are not covered in the Tour)

package main

import "fmt"

func sum(nums ...int) int {
	// TODO: Accumulate the numbers in nums and return the total.
}

func main() {
	fmt.Println(sum(1, 2, 3))
}
