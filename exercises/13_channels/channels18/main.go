// Problem: a caller starts several workers and must know when all of them have
// stopped after one cancellation broadcast.
//
// Without this pattern: one worker's completion signal cannot prove that every
// worker exited, so the caller may return while work is still running.
//
// Channels: stop is caller-owned and receive-only here. Closing it broadcasts
// cancellation. Each worker sends one token on exited; the coordinator closes
// done after it receives count tokens.
//
// Timeline:
//   start count workers
//   close(stop)
//   each worker exits and sends one exited token
//   coordinator receives count tokens and closes(done)
//
// Hint:
//   In the loop, start each worker with:
//
//       go func() {
//           <-stop
//           onWorkerExit()
//           exited <- struct{}{}
//       }()
//
//   Use a buffered exited channel so every worker can report without depending
//   on the coordinator's exact scheduling. A coordinator receives exactly count
//   tokens and closes done once. For count <= 0, return an already-closed done.

package main

import "fmt"

var onWorkerExit = func() {}

func startWorkers(count int, stop <-chan struct{}) <-chan struct{} {
	return nil // TODO: start workers, collect their exit signals, and close done once
}

func main() {
	stop := make(chan struct{})
	done := startWorkers(3, stop)
	close(stop)
	<-done
	fmt.Println("workers stopped")
}
