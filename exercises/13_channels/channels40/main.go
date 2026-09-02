// Problem: a relay is a reusable pipeline stage, so either its upstream or its
// downstream may become unavailable first.
// Without this pattern: a blocked receive or send leaves the stage running after
// the rest of the pipeline has stopped.
// Channels: in is receive-only upstream data; stop is cancellation; out is
// producer-owned output and is closed by the relay.
// Timeline:
//   receive in or stop
//   check comma-ok
//   send out or stop
//   repeat
//   close(out)
//
// Hint:
//   Defer close(out).
//   Select between stop and receiving in.
//   Check comma-ok.
//   Select between stop and sending the value to out.
//   This repeats the two-sided cancellation protocol from the earlier
//   relay exercises.

package main

import "fmt"

func relay(stop <-chan struct{}, in <-chan int) <-chan int {
	return nil // TODO: make both relay directions cancellation-aware
}
func main() { fmt.Println(relay(make(chan struct{}), nil)) }
