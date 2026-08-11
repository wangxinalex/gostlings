// Concept: result and done are different signals
// Task: publish one result, close the result channel, then close the caller's done channel
// Expected behavior: callers receive the result data separately from completion notification
// Hint: make a capacity-one result channel. In one goroutine, defer close(done), send 42,
//       and close the result channel before the goroutine returns. Never send data on done.

package main

import "fmt"

func runWithDone(done chan struct{}) <-chan int {
	return nil // TODO: use separate result and done channel lifecycles
}

func main() {
	done := make(chan struct{})
	for value := range runWithDone(done) {
		fmt.Println(value)
	}
	<-done
}
