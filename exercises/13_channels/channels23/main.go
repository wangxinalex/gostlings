// Problem: a fan-in consumer may abandon out while one input remains silent or
// while a forwarder is blocked sending to out.
//
// Without this pattern: a forwarder leaks at whichever blocking operation is not
// cancellation-aware, and the coordinator can never close out.
//
// Channels: stop is a receive-only cancellation broadcast. inputs carry data.
// out is shared and coordinator-owned. exited confirms forwarder termination.
//
// Timeline:
//   select: input or stop
//   select: out or stop
//   forwarder exits -> coordinator closes out
//
// Hint:
//   Protect both sides. Use a select with `<-stop` beside each input receive,
//   then a second select with `<-stop` beside `out <- value`. Each forwarder
//   sends one exit token on every return path; only the coordinator closes out.

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
