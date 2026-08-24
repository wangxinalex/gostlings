// Problem: the caller should receive a result channel immediately while work
// continues independently.
// Without this pattern: an unbuffered result send can keep a finished worker
// blocked until the caller happens to receive.
// Channels: out carries exactly one result and has capacity one; the worker
// owns the send and close if a stream close is added later.
// Timeline: return out -> work finishes -> result enters buffer -> caller receives
// Hint: make out with capacity 1 before starting the goroutine. The buffer lets
// the worker publish its one result even when the caller is briefly late.

package main

import "fmt"

func runAsync(work func() int) <-chan int {
	return nil // TODO: return a capacity-one result channel and run work in a goroutine
}

func main() {
	fmt.Println(<-runAsync(func() int { return 42 }))
}
