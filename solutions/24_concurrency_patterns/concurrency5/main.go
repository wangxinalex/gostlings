// Concept: a worker pool needs a bounded worker count, cancellation, and a final join.
// Task: process jobs with at most limit workers and return sorted results or an error.
// Hint: select on ctx.Done while receiving jobs; wait for workers before returning.
package main

import (
	"context"
	"errors"
	"sort"
	"sync"
)

var poolWork = func(ctx context.Context, job int) (int, error) {
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
		return job * 2, nil
	}
}

func runPool(ctx context.Context, workers, limit int, jobs []int) ([]int, error) {
	if workers < 1 || limit < 1 {
		return nil, errors.New("workers and limit must be positive")
	}
	if workers > limit {
		workers = limit
	}
	jobCh := make(chan int)
	results := make(chan int, len(jobs))
	errorsCh := make(chan error, 1)
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
					value, err := poolWork(ctx, job)
					if err != nil {
						select {
						case errorsCh <- err:
						default:
						}
						return
					}
					results <- value
				}
			}
		}()
	}
	for _, job := range jobs {
		select {
		case jobCh <- job:
		case <-ctx.Done():
			close(jobCh)
			wg.Wait()
			return nil, ctx.Err()
		}
	}
	close(jobCh)
	wg.Wait()
	select {
	case err := <-errorsCh:
		return nil, err
	default:
	}
	close(results)
	values := make([]int, 0, len(results))
	for value := range results {
		values = append(values, value)
	}
	sort.Ints(values)
	return values, nil
}

func main() {}
