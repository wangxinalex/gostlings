// Concept: slices share an underlying array; copy for independence
// Task: fix this program so mutating b leaves a untouched and it prints [1 2 3]
// Expected output: [1 2 3]
// Hint: assigning a slice copies only the view, not the data; use make and copy to duplicate it (Go Tour: Moretypes 8)

package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := make([]int, len(a))
	copy(b, a)
	b[0] = 100
	fmt.Println(a)
}
