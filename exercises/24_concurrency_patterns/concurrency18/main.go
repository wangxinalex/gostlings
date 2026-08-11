// Concept: a service combines request/reply, bounded workers, ordered responses,
// first-error cancellation, metrics, and graceful shutdown.
// Task: implement the request service without returning before all workers stop.
// Hint: use one derived context, a bounded semaphore, indexed responses, and a final join.
package main

import (
	"context"
	"errors"
)

type request struct {
	value int
	fail  bool
}
type response struct {
	value int
	err   error
}

var errRequestFailed = errors.New("request failed")

func runService(ctx context.Context, workers, limit int, requests []request) ([]response, error) {
	// TODO: Bound work, preserve response order, cancel on first error, and join workers.
	if workers < 1 || limit < 1 {
		return nil, errors.New("workers and limit must be positive")
	}
	return nil, nil
}
