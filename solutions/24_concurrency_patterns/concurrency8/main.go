// Concept: rate-limited work selects between token acquisition and cancellation.
// Task: consume one token per job and return processed results in input order.
// Hint: never wait for a token without a ctx.Done case.
package main

import (
	"context"
	"errors"
)

func runRateLimited(ctx context.Context, tokens <-chan struct{}, workers int, jobs []int) ([]int, error) {
	if workers < 1 {
		return nil, errors.New("workers must be positive")
	}
	results := make([]int, 0, len(jobs))
	for _, job := range jobs {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-tokens:
			results = append(results, job*2)
		}
	}
	return results, nil
}

func main() {}
