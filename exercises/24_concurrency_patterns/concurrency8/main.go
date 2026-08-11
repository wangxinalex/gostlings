// Concept: rate-limited work selects between token acquisition and cancellation.
// Task: consume one token per job and return processed results in input order.
// Hint: never wait for a token without a ctx.Done case.
package main

import (
	"context"
	"errors"
)

func runRateLimited(ctx context.Context, tokens <-chan struct{}, workers int, jobs []int) ([]int, error) {
	// TODO: Acquire one token per job with cancellation and process every job.
	if workers < 1 {
		return nil, errors.New("workers must be positive")
	}
	return nil, nil
}
