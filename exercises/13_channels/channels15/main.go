// Concept: cancellable fan-in — downstream abandonment must stop forwarders
// Task: merge inputs, but stop receiving and sending when stop is closed
// Expected behavior: a blocked input cannot keep the merged output alive after cancellation
// Hint: select on stop while receiving from each input and while sending to out

package main

import "fmt"

func merge(stop <-chan struct{}, inputs ...<-chan int) <-chan int {
	// Thought: responding to stop only around out <- value is not enough. A
	// forwarder may block on <-input, so both receive and send must be cancellable.
	return nil // TODO: add cancellable forwarders and close out after they exit
}

func main() {
	input := make(chan int, 2)
	input <- 1
	input <- 2
	close(input)
	for value := range merge(make(chan struct{}), input) {
		fmt.Println(value)
	}
}
