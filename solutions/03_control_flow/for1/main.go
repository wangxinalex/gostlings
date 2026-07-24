// Concept: the for loop
// Task: add a for loop that adds the numbers 1 through 10 to sum
// Expected output: 55
// Hint: for is Go's only loop; it has an init, a condition, and a post statement (Go Tour: Flowcontrol 1)

package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
