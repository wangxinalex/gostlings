// Concept: a service combines request/reply, bounded workers, ordered responses,
// first-error cancellation, metrics, and graceful shutdown.
// Task: implement the request service without returning before all workers stop.
// Hint: use one derived context, a bounded semaphore, indexed responses, and a final join.
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
			if req.fail {
				mu.Lock()
				if firstErr == nil {
					firstErr = errRequestFailed
					cancel()
				}
				mu.Unlock()
				return
			}
			responses[index] = response{value: req.value * 2}
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
