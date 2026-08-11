// Concept: relay — cancellation must cover both receiving and sending
// Task: forward values from in until in closes or stop closes, then close the output
// Expected behavior: normal input values arrive in order; cancellation ends a blocked send
// Hint: start a goroutine and defer close(out). Use one select to receive from in or
//       stop, then another select to send to out or stop. Check the comma-ok result.

package main

import "fmt"

func relay(stop <-chan struct{}, in <-chan int) <-chan int {
	return nil // TODO: make out and relay each receive/send with stop cases
}

func main() {
	in := make(chan int, 2)
	in <- 4
	in <- 9
	close(in)
	for value := range relay(make(chan struct{}), in) {
		fmt.Println(value)
	}
}
