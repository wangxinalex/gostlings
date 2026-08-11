// Concept: ordered collection separates execution order from response order.
// Task: process indexed jobs and return results in input order until cancellation or failure.
// Hint: store by index; use errors.Is for the documented failure.
package main

import (
	"context"
	"errors"
)

type orderedJob struct {
	value int
	fail  bool
}

var errOrderedJob = errors.New("ordered job failed")

func ordered(ctx context.Context, workers int, jobs []orderedJob) ([]int, error) {
	if workers < 1 {
		return nil, errors.New("workers must be positive")
	}
	results := make([]int, len(jobs))
	for i, job := range jobs {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		if job.fail {
			return nil, errOrderedJob
		}
		results[i] = job.value * 2
	}
	return results, nil
}

func main() {}
