// Concept: select with default can make a send non-blocking
// Task: send value when ch is ready, otherwise return immediately without sending
// Expected behavior: a ready send returns true; a full buffer or unready receiver returns false
// Hint: use select with case ch <- value: return true and a default case that returns false

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
