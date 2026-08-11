// Concept: draining a closed channel with comma-ok receives
// Task: collect every buffered value, then stop only when the receive reports that the channel is closed
// Expected behavior: values sent before close are returned in send order, including a real zero value
// Hint: use value, ok := <-ch in a loop; append while ok is true and return only when ok is false

package main

import "fmt"

func drainClosed(ch <-chan int) []int {
	// Thought: a closed channel can still have buffered values. Check ok after
	// every receive so a real zero value is not mistaken for the closed state.
	return nil // TODO: receive until ok is false, collecting every value first
}

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	fmt.Println(drainClosed(ch))
}
