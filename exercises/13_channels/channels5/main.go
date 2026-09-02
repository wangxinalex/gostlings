// Problem: a channel may be closed while buffered data is still waiting.
//
// Without this pattern: stopping on close or on value == 0 can lose data or
// reject a legitimate zero value.
//
// Channels: ch is a receive-only data stream. The sender already closed it.
//
// Timeline:
//   buffered values -> close(ch) -> receive all values -> ok=false
//
// Hint:
//   Receive with comma-ok in a loop. Append while ok is true, and return only
//   after the receive reports that ch is closed and drained.

package main

import "fmt"

func drainClosed(ch <-chan int) []int {
	// Thought: a closed channel can still have buffered values. Check ok after
	// every receive so a real zero value is not mistaken for the closed state.
	return nil // TODO: receive until ok is false, collecting every value first
}

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	fmt.Println(drainClosed(ch))
}
