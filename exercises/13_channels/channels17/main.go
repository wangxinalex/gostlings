// Concept: cancellable forwarder — both sides of a relay can block
// Task: forward input values until input closes or stop closes, then close the output
// Expected behavior: stop releases a blocked receive and a blocked send
// Hint: use a receive select with cases for in and stop. After receiving a value, use
//       another select with cases for out <- value and stop. Defer close(out) once.

package main

import "fmt"

func forward(stop <-chan struct{}, in <-chan int) <-chan int {
	return nil // TODO: make both the receive and send cancellation-aware
}

func main() {
	in := make(chan int, 2)
	in <- 3
	in <- 8
	close(in)
	for value := range forward(make(chan struct{}), in) {
		fmt.Println(value)
	}
}
