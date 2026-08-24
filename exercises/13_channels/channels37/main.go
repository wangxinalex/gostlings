// Problem: downstream may stop reading while upstream still has work to send.
// Without this pattern: the upstream forwarder remains blocked on out and leaks.
// Channels: work carries upstream data; stop is the downstream cancellation
// request; out is owned and closed by this forwarder.
// Timeline: receive work/stop -> send out/stop -> repeat -> close(out)
// Hint: defer close(out). Use one select for the work receive and another for
// the out send, with a stop case in both. Only this forwarder closes out; the
// caller closes stop when it abandons the stream.
package main

import "fmt"

func collectOrStop(stop <-chan struct{}, work <-chan int) <-chan int {
	return nil // TODO: make the receive and send cancellable
}
func main() { fmt.Println(collectOrStop(make(chan struct{}), nil)) }
