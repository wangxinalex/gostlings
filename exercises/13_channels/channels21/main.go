// Concept: cancellation-aware pipelines stop both receives and sends
// Task: apply two stages while allowing stop to interrupt a blocked input or output
// Expected behavior: closing stop closes the final output and lets every stage exit
// Hint: select on stop when receiving from input and sending to output in every stage

package main

import "fmt"

func pipeline(stop <-chan struct{}, in <-chan int) <-chan int {
	// Thought: a pipeline is a chain of goroutines. Cancelling only the final
	// stage can leave upstream stages blocked, so the same stop signal must reach
	// every stage's receive and send operations.
	return nil // TODO: compose cancellable stages
}

func main() {
	in := make(chan int, 2)
	in <- 1
	in <- 2
	close(in)
	for value := range pipeline(make(chan struct{}), in) {
		fmt.Println(value)
	}
}
