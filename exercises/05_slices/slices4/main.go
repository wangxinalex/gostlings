// Concept: copying slices with make and copy
// Task: use make and copy so mutating b leaves a untouched and it prints [1 2 3]
// Expected output: [1 2 3]
// Hint: make([]int, len(a)) allocates a slice with its own backing array;
//       copy(dst, src) copies values between slices (Go Tour: Moretypes 8)

package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := make([]int, len(a))
	copy(b, nil) // TODO: Replace nil with a.
	b[0] = 100
	fmt.Println(a)
}
