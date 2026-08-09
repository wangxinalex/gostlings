// Concept: variadic functions and range
// Task: complete sum so it adds up all its arguments
// Expected output: 6
// Hint: a variadic parameter collects arguments into a slice. The for range
//       loop below introduces how to visit each value; add each n to total.
//       (variadic parameters are not covered in the Go Tour)

package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		// TODO: Add n to total.
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3))
}
