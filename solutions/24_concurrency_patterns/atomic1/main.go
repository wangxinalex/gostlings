// Concept: atomic counters and WaitGroup coordination
// Task: increment a shared counter from many goroutines without a data race
// Expected output: focused race-aware test passes (run `go test -race ./solutions/24_concurrency_patterns/atomic1`)
// Hint: use atomic.Int64.Add for each increment, call wg.Done in every worker, and wait before Load

package main

import (
	"sync"
	"sync/atomic"
)

func incrementConcurrently(workers, increments int) int64 {
	var counter atomic.Int64
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < increments; j++ {
				counter.Add(1)
			}
		}()
	}

	wg.Wait()
	return counter.Load()
}
