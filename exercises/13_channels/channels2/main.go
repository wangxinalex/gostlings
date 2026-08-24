// Problem: a producer wants to publish two values before the consumer reads.
// Without this pattern: an unbuffered send waits immediately for a receiver.
// Channels: ch carries data; its buffer is the temporary handoff capacity.
// Timeline: send 1 -> buffer; send 2 -> buffer; receive 1; receive 2
// Hint: give ch capacity for exactly two values. A buffer delays blocking; it
// does not make sends unlimited or remove the need for a receiver (Go Tour: Concurrency 3).

package main

import "fmt"

func main() {
	// Thought: a buffer of two can hold two values so sends can complete first;
	// once the buffer is full, another send still blocks.
	// Pattern:
	//   ch := make(chan T, capacity)
	//   send up to capacity values without a receiver
	//   receive values later; sending blocks when the buffer is full
	ch := make(chan int) // TODO: Make this buffered so the sends don't block.

	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
