// Concept: growing a slice with append
// Task: fix this program so it compiles and prints [1 2 3]
// Expected output: [1 2 3]
// Hint: append returns the extended slice; assign the result back to s (Go Tour: Moretypes 15)

package main

import "fmt"

func main() {
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Println(s)
}
