// Concept: building a slice of slices
// Task: complete chunk so it splits values into consecutive groups of at most
//       size elements; the last group may be shorter and size < 1 returns nil
// Expected output: [[1 2 3] [4 5 6] [7]]
// [[1 2] [3 4] [5 6] [7]]
// size 0 is nil: true
// Hint: start at 0 and step by size; clamp end to len(values), then append
//       values[start:end] to a [][]int. The last line prints true only when
//       chunk(values, 0) really returns nil (Go Tour: Moretypes 11-14)

package main

import "fmt"

func chunk(values []int, size int) [][]int {
	// TODO: Build a [][]int of consecutive groups of at most size elements.
	return nil
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(chunk(values, 3))
	fmt.Println(chunk(values, 2))
	fmt.Println("size 0 is nil:", chunk(values, 0) == nil)
}
