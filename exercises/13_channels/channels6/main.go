// Concept: select with a default case — non-blocking channel operations
// Task: the channel is empty (no goroutine ever sends); add a default case so we don't block forever
// Expected output: no value available, doing other work
// Hint: select { case v := <-ch: ... default: ... } — the default branch runs immediately when no channel is ready (Go Tour: Concurrency 5-6)

package main

import "fmt"

func main() {
	ch := make(chan int)

	// No one ever sends into ch, so <-ch would block forever.
	// TODO: Use a select with a default case to handle the empty channel gracefully.

	// The line below blocks — replace it with the select.
	val := <-ch
	fmt.Println("received:", val)
}
