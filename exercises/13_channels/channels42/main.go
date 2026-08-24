// High-frequency reinforcement: cancellation must protect producer sends.
// Problem: a consumer may read only part of a stream and abandon the rest.
// Without this pattern: the producer blocks forever on the next output send.
// Channels: stop is a receive-only cancellation request; out carries values and
// is closed by the producer on normal completion or cancellation.
// Timeline: produce value -> select out/stop -> repeat -> close(out)
// Hint: create out and defer close(out) in one goroutine. For every value, select
// between `out <- value` and `<-stop`; never leave a producer blocked on an
// abandoned output send.
package main

import "fmt"

func produce(stop <-chan struct{}) <-chan int {
	return nil // TODO: make each output send stop-aware and close out on every exit path
}

func main() { fmt.Println(produce(make(chan struct{}))) }
