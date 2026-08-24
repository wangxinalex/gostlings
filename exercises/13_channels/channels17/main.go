// Problem: a forwarder can be stuck before it has a value, or after it has a
// value but no downstream receiver is ready.
// Without this pattern: cancellation may release one wait while leaving the
// goroutine blocked on the other side.
// Channels: in carries caller-owned data; stop requests cancellation; out is
// owned and closed by the forwarder.
// Timeline: select receive in/stop -> select send out/stop -> repeat -> close(out)
// Hint: use comma-ok for the input receive and two separate selects: one around
// receiving, one around sending. Defer close(out) exactly once.

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
