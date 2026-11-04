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
	"sync"
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
	if workers < 1 || limit < 1 {
		return nil, errors.New("workers and limit must be positive")
	}
	serviceCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	capacity := workers
	if limit < capacity {
		capacity = limit
	}
	sem := make(chan struct{}, capacity)
	responses := make([]response, len(requests))
	var wg sync.WaitGroup
	var mu sync.Mutex
	var firstErr error
	for index, req := range requests {
		select {
		case <-serviceCtx.Done():
			break
		case sem <- struct{}{}:
		}
		if serviceCtx.Err() != nil {
			break
		}
		wg.Add(1)
		go func(index int, req request) {
			defer wg.Done()
			defer func() { <-sem }()
			select {
			case <-serviceCtx.Done():
				return
			default:
			}
			value, err := serviceWork(serviceCtx, req)
			if err != nil {
				mu.Lock()
				if firstErr == nil {
					firstErr = err
					cancel()
				}
				mu.Unlock()
				return
			}
			responses[index] = response{value: value}
		}(index, req)
	}
	wg.Wait()
	if firstErr != nil {
		return nil, firstErr
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return responses, nil
}

func main() {}
