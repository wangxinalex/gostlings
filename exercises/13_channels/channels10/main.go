// Problem: a producer wants to publish opportunistically without waiting for a
// receiver or free buffer space.
// Without this pattern: a send can block the producer indefinitely.
// Channels: ch is a send-only data channel; default means this attempt will not
// be queued when the send is not ready.
// Timeline: try send -> ready: value enters ch; not ready: default -> return false
// Hint: select between `ch <- value` and default. Return true only from the send
// case and false from default.

package main

import "fmt"

func trySend(ch chan<- int, value int) bool {
	// Thought: default means this is one attempt. It must not be placed in a
	// loop that retries continuously without useful work or a backoff.
	return false // TODO: select between sending value and returning immediately
}

func main() {
	ch := make(chan int, 1)
	fmt.Println(trySend(ch, 7))
}
