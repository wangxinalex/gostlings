// Problem: a caller needs to wait for asynchronous work without receiving a
// meaningless payload.
//
// Without this pattern: every waiter would need a separate result or a guessed
// value, and multiple waiters could miss a one-time send.
//
// Channels: done carries no data. The completing goroutine owns close(done).
//
// Timeline:
//   start work -> work finishes -> close(done) -> every waiter proceeds
//
// Hint:
//   Make done, start one goroutine, and defer close(done) at the top of that
//   goroutine. Closing broadcasts completion to all receivers; do not send a value.

package main

import "fmt"

func complete() <-chan struct{} {
	return nil // TODO: close a done channel from the completing goroutine
}

func main() {
	<-complete()
	fmt.Println("complete")
}
