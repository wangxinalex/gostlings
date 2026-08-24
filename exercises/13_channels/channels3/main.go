// Problem: the consumer uses range, so it needs a completion signal after the
// final value.
// Without this pattern: range keeps waiting for another value after 3 and the
// program deadlocks.
// Channels: ch carries data and its close means “no more sends”; the sender
// owns the close because it knows when the stream is complete.
// Timeline: send 1,2,3 -> close(ch) -> range drains values -> range stops
// Hint: close ch after the final send. Closing does not discard already-sent
// values; range receives them before it ends (Go Tour: Concurrency 4).

package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		// Thought: range needs to know that no more values will arrive. The sender
		// closes the channel after the final send, so range drains existing values
		// before it exits.
		// TODO: Close the channel so the range below does not deadlock.
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
