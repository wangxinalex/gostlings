// Problem: a caller gives up after a deadline, but the producer may still be
// sleeping or blocked on its result send.
//
// Without this pattern: returning on timeout alone leaks the producer and may
// leave it using resources after its caller has gone away.
//
// Channels: result carries business data. stop carries the cancellation request.
// done carries the producer's exit confirmation. run owns close(stop); the
// wrapper goroutine owns close(done).
//
// Timeline:
//   start producer
//   wait for result or 25ms timeout
//   close(stop)
//   wait for <-done
//   return
//
// Hint:
//   Create `stop` and a capacity-one `result`, then run the producer in a
//   goroutine that defers closing `done`. Wait for either the result or the
//   deadline; in both branches close stop, wait for done, and only then return.
//   The timeout branch returns exactly "timed out".
// Stuck?: time.After supplies the 25ms deadline case; the buffer lets a late
//   producer publish without a receiver.

package main

import "fmt"

var runProducer = func(stop <-chan struct{}, result chan<- string) {}

func run(done chan struct{}) string {
	return "" // TODO: cancel the producer on timeout and join it before returning
}

func main() {
	done := make(chan struct{})
	fmt.Println(run(done))
}
