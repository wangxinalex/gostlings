// Concept: a request combines deadline propagation, cooperative workers, and a final join.
// Task: run workers under a timeout, return one summary on success, and return the
// context error only after every worker has stopped.
// Hint: derive one timeout context, select on ctx.Done in every worker, and join
// workers through an acknowledgment channel before returning.
package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var withRequestTimeout = context.WithTimeout
var requestWork = func(ctx context.Context, id int) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
		return fmt.Sprintf("request-%d", id), nil
	}
}

func runRequest(ctx context.Context, workers int) (string, error) {
	if workers < 1 {
		return "", errors.New("workers must be positive")
	}
	requestCtx, cancel := withRequestTimeout(ctx, time.Second)
	defer cancel()

	type result struct{ err error }
	results := make(chan result, workers)
	exited := make(chan struct{}, workers)
	for id := 0; id < workers; id++ {
		go func(id int) {
			_, err := requestWork(requestCtx, id)
			results <- result{err: err}
			exited <- struct{}{}
		}(id)
	}
	joined := make(chan struct{})
	go func() {
		for range workers {
			<-exited
		}
		close(joined)
	}()

	<-joined
	for range workers {
		if err := (<-results).err; err != nil {
			return "", err
		}
	}
	return fmt.Sprintf("processed %d requests", workers), nil
}

func main() {}
