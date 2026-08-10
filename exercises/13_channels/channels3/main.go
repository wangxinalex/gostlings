// Concept: closing a channel tells range loops to stop
// Task: this program deadlocks because range never stops; close the channel after sending
// Expected output: 1
// 2
// 3
// Hint: close(ch) after the sends so the range loop knows there are no more values (Go Tour: Concurrency 4)

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
