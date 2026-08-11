// Concept: a worker pool needs a bounded worker count, cancellation, and a final join.
// Task: process jobs with at most limit workers and return sorted results or an error.
// Hint: select on ctx.Done while receiving jobs; wait for workers before returning.
package main

import (
	"context"
	"errors"
)

var poolWork = func(ctx context.Context, job int) (int, error) {
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
		return job * 2, nil
	}
}

func runPool(ctx context.Context, workers, limit int, jobs []int) ([]int, error) {
	// TODO: Bound active workers by limit, cancel blocked work, and join every worker.
	if workers < 1 || limit < 1 {
		return nil, errors.New("workers and limit must be positive")
	}
	return nil, nil
}
