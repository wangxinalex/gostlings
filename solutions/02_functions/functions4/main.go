// Concept: variadic functions
// Task: complete the body of sum so it adds up all its arguments
// Expected output: 6
// Hint: variadic parameters collect arguments into a slice; use range to iterate over it

package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3))
}
