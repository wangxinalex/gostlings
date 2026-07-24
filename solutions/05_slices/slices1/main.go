// Concept: creating a slice with a slice literal
// Task: create an int slice containing 1, 2, 3 so the program prints its length
// Expected output: length: 3
// Hint: a slice literal looks like []int{1, 2, 3} (Go Tour: Moretypes 9)

package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Println("length:", len(s))
}
