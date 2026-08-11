// Concept: cancellation-aware workers need a final join signal.
// Task: start count workers, let each observe cancellation, and close done only after all return.
// Hint: use an acknowledgment channel and one coordinator; count zero must close done too.
package main

import "context"

var workerStarted = func() {}
var workerStopped = func() {}

func runWorkers(ctx context.Context, count int) <-chan struct{} {
	// TODO: Start workers that stop on ctx.Done, then close done after every worker exits.
	done := make(chan struct{})
	close(done)
	return done
}

func main() {}
