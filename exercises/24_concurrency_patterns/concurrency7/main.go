// Concept: batching combines a size threshold, explicit flush events, and cancellation.
// Task: emit full batches immediately and partial batches on flush; discard on cancel.
// Hint: keep one owner for the current slice and close output on every exit.
package main

import (
	"context"
	"time"
)

func batch(ctx context.Context, in <-chan int, flush <-chan time.Time, size int) <-chan []int {
	// TODO: Emit full and timer-flushed batches, then close output on cancellation.
	return nil
}
