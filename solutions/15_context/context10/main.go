// Concept: cancellation-aware workers need a final join signal.
// Task: start count workers, let each observe cancellation, and close done only after all return.
// Hint: use an acknowledgment channel and one coordinator; count zero must close done too.
package main

import "context"

var workerStarted = func() {}
var workerStopped = func() {}

func runWorkers(ctx context.Context, count int) <-chan struct{} {
	done := make(chan struct{})
	exited := make(chan struct{}, count)
	for range count {
		go func() {
			workerStarted()
			<-ctx.Done()
			workerStopped()
			exited <- struct{}{}
		}()
	}
	go func() {
		for range count {
			<-exited
		}
		close(done)
	}()
	return done
}

func main() {}
