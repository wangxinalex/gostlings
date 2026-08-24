// Problem: shutdown itself must initiate cancellation and return a waitable
// completion signal while workers finish asynchronously.
// Without this pattern: callers either close stop in several places and risk a
// double-close panic, or return before every worker has released its resources.
// Channels: stop is the shared cancellation broadcast; exited carries one
// acknowledgement per worker; done is the final completion broadcast. The
// coordinator is the only owner allowed to close stop and done.
// Timeline: start workers -> coordinator closes stop once -> workers send exited -> coordinator closes done
// Hint: create `exited := make(chan struct{}, workers)` and `done`. Start each
// worker in a loop: wait for `<-stop`, call onShutdownWorkerExit, then send one
// exited token. Start a coordinator goroutine that checks whether stop is
// already closed before closing it, receives exactly workers tokens, and closes
// done once. Return done immediately; the caller performs the join with `<-done`.

package main

import "fmt"

var onShutdownWorkerExit = func() {}

func shutdown(stop chan struct{}, workers int) <-chan struct{} {
	return nil // TODO: coordinate stop, worker exit signals, and one done close
}

func main() {
	stop := make(chan struct{})
	<-shutdown(stop, 3)
	fmt.Println("shutdown complete")
}
