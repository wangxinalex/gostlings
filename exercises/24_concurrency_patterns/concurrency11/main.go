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
	// TODO: Process jobs under ctx and return results in original order.
	if workers < 1 {
		return nil, errors.New("workers must be positive")
	}
	return nil, nil
}
