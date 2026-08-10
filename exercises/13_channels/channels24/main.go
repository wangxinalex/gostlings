// Concept: nil channels disable select cases
// Task: drain two inputs and set each one to nil after it closes
// Expected behavior: all values arrive once, with no repeated zero values after close
// Hint: in a select loop, assign a closed input variable to nil to remove that case

package main

import "fmt"

func drain(first, second <-chan int) []int {
	// Thought: a closed channel is always readable and keeps producing zero
	// values. Set the corresponding variable to nil to disable that select case
	// permanently after the input is drained.
	return nil // TODO: disable each input after its close and collect values
}

func main() {
	first := make(chan int, 2)
	second := make(chan int, 2)
	first <- 1
	first <- 3
	second <- 2
	second <- 4
	close(first)
	close(second)
	fmt.Println(drain(first, second))
}
