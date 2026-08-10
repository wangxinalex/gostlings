// Concept: a pipeline stage owns and closes its output
// Task: square every input value and close the output after input closes
// Expected behavior: the stage can be ranged over until completion
// Hint: one goroutine ranges in, sends transformed values, and defers close(out)

package main

import "fmt"

func square(in <-chan int) <-chan int {
	// Thought: each stage closes only its own output. It never closes an input
	// channel supplied by its caller.
	return nil // TODO: transform values in a goroutine and close the output
}

func main() {
	in := make(chan int, 3)
	in <- 1
	in <- 2
	in <- 3
	close(in)
	for value := range square(in) {
		fmt.Println(value)
	}
}
