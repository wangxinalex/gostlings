// Problem: an unbuffered channel is a synchronous handoff, but this program
// sends and receives in the same goroutine.
// Without this pattern: the send waits for a receiver that this goroutine has
// not reached yet, so the program deadlocks.
// Channels: ch carries one string; main owns both sides for this exercise.
// Timeline: sender -- waits for receiver --> ch --> receiver
// Hint: move the send into a goroutine. The receiver in main can then rendezvous
// with it (Go Tour: Concurrency 2).

package main

import "fmt"

func main() {
	ch := make(chan string)

	// Thought: sending on an unbuffered channel must rendezvous with a receiver.
	// TODO: Move the send into a goroutine so this doesn't deadlock.
	ch <- "hi"

	fmt.Println(<-ch)
}
