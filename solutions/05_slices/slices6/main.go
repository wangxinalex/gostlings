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
	if len(values) == 0 {
		return values
	}
	kept := values[:0]
	kept = append(kept, values[0])
	for _, v := range values[1:] {
		if v != kept[len(kept)-1] {
			kept = append(kept, v)
		}
	}
	return kept
}

func main() {
	values := []int{1, 1, 2, 3, 3, 3, 4, 1}
	compacted := compact(values)
	fmt.Println(compacted)
	fmt.Println("same backing array:", &compacted[0] == &values[0])
}
