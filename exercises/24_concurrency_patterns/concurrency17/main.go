// Concept: shutdown joins in-flight work before returning an error.
// Task: process jobs under cancellation and return collected results only after workers stop.
// Hint: use a result channel, a worker join, and errors.Is-compatible context errors.
package main

import (
	"context"
	"errors"
)

type shutdownJob struct {
	value int
	fail  bool
}
type shutdownResult struct {
	value int
	err   error
}

func shutdown(ctx context.Context, workers int, jobs []shutdownJob) ([]shutdownResult, error) {
	// TODO: Run jobs, join workers, and return a cancellation or job error after cleanup.
	if workers < 1 {
		return nil, errors.New("workers must be positive")
	}
	return nil, nil
}
