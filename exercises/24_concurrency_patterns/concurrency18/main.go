// Concept: a service combines request/reply, bounded workers, ordered responses,
// and first-error cancellation, and returns only after its workers have stopped.
// Task: implement the request service through serviceWork without returning early.
// Expected behavior: responses keep request order; no more than min(workers, limit)
// requests run at once; the first failing request cancels the rest; every started
// request finishes before the service returns; a canceled caller is reported as
// context.Canceled.
// Observability: handle each request through serviceWork. Tests replace it to
// count active work and to hold workers open while checking the final join.
// Hint: derive one context, bound workers with a semaphore, index each
// response, cancel on the first error, and join before returning.
// Stuck?: context.WithCancel, sync.WaitGroup
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

var serviceWork = func(ctx context.Context, req request) (int, error) {
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}
	if req.fail {
		return 0, errRequestFailed
	}
	return req.value * 2, nil
}

func runService(ctx context.Context, workers, limit int, requests []request) ([]response, error) {
	// TODO: Bound work, preserve response order, cancel on first error, and join workers.
	if workers < 1 || limit < 1 {
		return nil, errors.New("workers and limit must be positive")
	}
	return nil, nil
}
