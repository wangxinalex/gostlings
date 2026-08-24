// Problem: a relay can block while waiting for input and again while delivering
// a value to a slow downstream consumer.
// Without this pattern: protecting only one side leaves a goroutine leaked on
// the other blocked operation.
// Channels: in carries data and is caller-owned; stop is a receive-only cancel
// signal; out carries data and is producer-owned and closed by the relay.
// Timeline: receive in or stop -> send out or stop -> repeat -> close(out)
// Hint: use one select for the receive and a second select for the send. Check
// comma-ok on in so normal input closure also ends the relay.

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
