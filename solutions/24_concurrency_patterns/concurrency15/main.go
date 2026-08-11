// Concept: bounded submission should shed load instead of blocking forever.
// Task: enqueue one job when capacity is available, or return a clear rejection/cancellation error.
// Hint: use a non-blocking send plus a ctx.Done case; result is reserved for the receiver protocol.
package main

import (
	"context"
	"errors"
)

type submitJob struct{ value int }
type submitResult struct {
	value    int
	accepted bool
}

var errQueueFull = errors.New("queue full")

func submit(ctx context.Context, queue chan<- submitJob, result <-chan submitResult, capacity int) error {
	_ = result
	job := submitJob{value: capacity}
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	select {
	case queue <- job:
		return nil
	default:
		return errQueueFull
	}
}

func main() {}
