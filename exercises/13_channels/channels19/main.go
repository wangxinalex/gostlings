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
//   Create `stop` and a capacity-one `result`, then wrap runProducer in a
//   goroutine with `defer close(done)`. Select on result versus
//   `<-time.After(25*time.Millisecond)`. In both branches close stop, wait for
//   done, and only then return. The timeout branch returns exactly "timed out".

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
