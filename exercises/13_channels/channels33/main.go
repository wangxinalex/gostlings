// Problem: several forwarders share one output, so completion can happen in any
// order but output closure must happen exactly once.
// Without this pattern: allowing each forwarder to close out causes a double
// close; closing out after one forwarder drops values from the others.
// Channels: sources are receive-only inputs; out is shared data owned by the
// collector; exited carries one acknowledgement per forwarder.
// Timeline:
//   sources drain independently
//   each forwarder sends one exit acknowledgement
//   coordinator receives all acknowledgements
//   coordinator closes out
//
// Hint:
//   Give each forwarder one exit acknowledgement.
//   A coordinator receives every acknowledgement.
//   The coordinator then closes out.
//   No forwarder may close the shared output.

package main

import "fmt"

var onCollectorExit = func() {}

func collect(sources []<-chan int) <-chan int {
	return nil // TODO: forward every source and coordinate close(out)
}
func main() { fmt.Println(collect(nil)) }
