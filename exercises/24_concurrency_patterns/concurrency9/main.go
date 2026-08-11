// Concept: graceful service shutdown stops intake before closing results.
// Task: process jobs until cancellation, close results, then close done.
// Hint: one goroutine owns both closures and checks ctx before each receive/send.
package main

import "context"

func serve(ctx context.Context, jobs <-chan int) (<-chan int, <-chan struct{}) {
	// TODO: Process jobs, close results, and signal done after shutdown.
	return nil, nil
}
