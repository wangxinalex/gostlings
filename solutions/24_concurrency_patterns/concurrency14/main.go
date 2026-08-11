// Concept: atomic metrics are read after workers join.
// Task: count completed and canceled jobs without a data race.
// Hint: use atomic counters and make each worker observe ctx before reporting its outcome.
package main

import (
	"context"
	"sync"
	"sync/atomic"
)

var measureWork = func(ctx context.Context, job int) bool { return ctx.Err() == nil }

func runMeasured(ctx context.Context, workers int, jobs []int) (completed, canceled int64) {
	if workers < 1 {
		workers = 1
	}
	var done sync.WaitGroup
	var completedCount atomic.Int64
	var canceledCount atomic.Int64
	jobCh := make(chan int)
	done.Add(workers)
	for range workers {
		go func() {
			defer done.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case job, ok := <-jobCh:
					if !ok {
						return
					}
					if measureWork(ctx, job) {
						completedCount.Add(1)
					} else {
						canceledCount.Add(1)
					}
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
	done.Wait()
	return completedCount.Load(), canceledCount.Load()
}

func main() {}
