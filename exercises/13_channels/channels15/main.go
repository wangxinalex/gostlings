// Problem: callers need both a business result and a reliable completion event.
// Without this pattern: using the result channel as completion makes it hard to
// distinguish “value is ready” from “the producer has fully exited”.
// Channels: result carries 42 and is closed by its producer; done carries no
// data and is closed by the completing goroutine.
// Timeline: send result -> close(result) -> close(done) -> caller may wait on either meaning
// Hint: make result capacity one. Send the value, close result, and defer close(done)
// in one goroutine. Never send business data on done.

package main

import "fmt"

func runWithDone(done chan struct{}) <-chan int {
	return nil // TODO: use separate result and done channel lifecycles
}

func main() {
	done := make(chan struct{})
	for value := range runWithDone(done) {
		fmt.Println(value)
	}
	<-done
}
