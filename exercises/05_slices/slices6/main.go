// Concept: filtering a slice in place
// Task: complete compact so it removes adjacent duplicates in one pass and
//       returns the shortened slice without allocating a new backing array
// Expected output: [1 2 3 4 1]
// same backing array: true
// Hint: keep := values[:0] shares values' backing array; append only the values
//       you keep and return keep. The second line prints true only when the
//       result still points at values[0] (Go Tour: Moretypes 7-15; Go doc: append)

package main

import "fmt"

func compact(values []int) []int {
	// TODO: Remove adjacent duplicates in one pass and return the shortened slice.
	return values
}

func main() {
	values := []int{1, 1, 2, 3, 3, 3, 4, 1}
	compacted := compact(values)
	fmt.Println(compacted)
	fmt.Println("same backing array:", &compacted[0] == &values[0])
}
