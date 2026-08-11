// Concept: shutdown joins in-flight work before returning an error.
// Task: process jobs under cancellation and return collected results only after workers stop.
// Hint: use a result channel, a worker join, and errors.Is-compatible context errors.
package main

import (
	"context"
	"errors"
	"sync"
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
	if workers < 1 {
		return nil, errors.New("workers must be positive")
	}
	if workers > len(jobs) && len(jobs) > 0 {
		workers = len(jobs)
	}
	if len(jobs) == 0 {
		return []shutdownResult{}, nil
	}
	jobCh := make(chan shutdownJob)
	results := make(chan shutdownResult, len(jobs))
	var wg sync.WaitGroup
	wg.Add(workers)
	for range workers {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case job, ok := <-jobCh:
					if !ok {
						return
					}
					if job.fail {
						results <- shutdownResult{err: errors.New("job failed")}
						return
					}
					results <- shutdownResult{value: job.value * 2}
				}
			}
		}()
	}
sendLoop:
	for _, job := range jobs {
		select {
		case jobCh <- job:
		case <-ctx.Done():
			break sendLoop
		}
	}
	close(jobCh)
	wg.Wait()
	close(results)
	var collected []shutdownResult
	for value := range results {
		collected = append(collected, value)
		if value.err != nil {
			return nil, value.err
		}
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return collected, nil
}

func main() {}
