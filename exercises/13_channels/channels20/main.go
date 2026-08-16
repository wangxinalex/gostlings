// Concept: graceful shutdown — stop, join workers, then report completion
// Task: close stop once, wait for workers, and close done exactly once
// Expected behavior: workers observe the stop broadcast before the coordinator closes done
// Hint: use this sequence:
//       create exited := make(chan struct{}, workers) and done
//       start each worker: <-stop, call onShutdownWorkerExit, then send one exited signal
//       start a coordinator goroutine that receives workers exited signals, then closes done once.
//       The coordinator owns close(stop): first check whether stop is already closed with a
//       non-blocking select, and only then close it. It must not close an already-closed stop.
//       shutdown must return done before waiting, so the caller can wait for completion separately.

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
