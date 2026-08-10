// Concept: a done channel reports completion without carrying data
// Task: close the done channel when the background operation finishes
// Expected behavior: complete() returns a channel that closes promptly
// Hint: the goroutine should defer close(done); a close is a broadcast notification

package main

import "fmt"

func complete() <-chan struct{} {
	done := make(chan struct{})
	// Thought: done carries only completion, not a result. Closing it wakes all
	// receivers waiting on it.
	return done // TODO: start work and close done when it finishes
}

func main() {
	<-complete()
	fmt.Println("completed")
}
