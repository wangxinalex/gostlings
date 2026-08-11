// Concept: cancellation must travel through every stage of a pipeline.
// Task: multiply each value by two, add one in a second stage, and close output on cancel.
// Hint: each stage needs cancellation-aware receive and send cases.
package main

import "context"

func pipeline(ctx context.Context, in <-chan int) <-chan int {
	// TODO: Build two context-aware stages and close the final output on every exit.
	return nil
}
