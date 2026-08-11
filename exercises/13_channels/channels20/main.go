// Concept: graceful shutdown — stop, join workers, then report completion
// Task: close stop once, wait for workers, and close done exactly once
// Expected behavior: workers observe the stop broadcast before the coordinator closes done
// Hint: workers wait on <-stop and send one exit signal to a buffered channel. The coordinator
//       owns close(stop) and close(done); first check whether stop is already closed before closing it.

package main

import "fmt"

func shutdown(stop chan struct{}, workers int) <-chan struct{} {
	return nil // TODO: coordinate stop, worker exit signals, and one done close
}

func main() {
	stop := make(chan struct{})
	<-shutdown(stop, 3)
	fmt.Println("shutdown complete")
}
