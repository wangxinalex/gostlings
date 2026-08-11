// Concept: a cancellable fan-in must stop on both sides of every forwarder.
// Task: merge inputs until they close or stop closes, then close out after all forwarders exit.
// Expected behavior: stop releases a forwarder blocked on either input receive or output send.
// Hint: use select with <-stop beside each input receive, then another select with <-stop beside out <- value.
//       Forwarders acknowledge exit; only the coordinator closes out.

package main

import "fmt"

var onMergeBeforeSend = func() {}

func merge(stop <-chan struct{}, inputs ...<-chan int) <-chan int {
	return nil // TODO: make both fan-in receives and sends cancellation-aware
}

func main() {
	in := make(chan int, 2)
	in <- 1
	in <- 2
	close(in)
	for value := range merge(make(chan struct{}), in) {
		fmt.Println(value)
	}
}
