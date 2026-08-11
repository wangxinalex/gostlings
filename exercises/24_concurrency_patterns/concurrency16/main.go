// Concept: buffering changes decoupling, not cancellation ownership.
// Task: double values through a buffered stage and close output on cancellation.
// Hint: make the buffer size explicit and retain cancellation cases in both directions.
package main

import "context"

func bufferedPipeline(ctx context.Context, in <-chan int, buffer int) <-chan int {
	// TODO: Build a buffered cancellation-aware stage and close output on all exits.
	return nil
}
