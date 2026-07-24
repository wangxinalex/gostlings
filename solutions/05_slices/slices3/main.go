// Concept: extracting a sub-slice with a slice expression
// Task: use a slice expression to extract [2 3 4] from s into sub
// Expected output: [2 3 4]
// Hint: s[low:high] selects elements low through high-1 (Go Tour: Moretypes 7)

package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	sub := s[1:4]
	fmt.Println(sub)
}
