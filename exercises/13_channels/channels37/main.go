// Problem: downstream may stop reading while upstream still has work to send.
// Without this pattern: the upstream forwarder remains blocked on out and leaks.
// Channels: work carries upstream data; stop is the downstream cancellation
// request; out is owned and closed by this forwarder.
// Timeline:
//   receive work or stop
//   send out or stop
//   repeat
//   close(out)
//
// Hint:
//   Defer close(out).
//   Use one select for the work receive.
//   Use another select for the out send.
//   Include a stop case in both selects.
//   Only this forwarder closes out.
//   The caller closes stop when it abandons the stream.

package main

import "fmt"

func collectOrStop(stop <-chan struct{}, work <-chan int) <-chan int {
	return nil // TODO: make the receive and send cancellable
}
func main() { fmt.Println(collectOrStop(make(chan struct{}), nil)) }
