// Concept: atomic counters and WaitGroup coordination
// Task: increment a shared counter from many goroutines without a data race
// Expected output: focused race-aware test passes (run `go test -race ./exercises/24_concurrency_patterns/atomic1`)
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

	// TODO: Start workers. Each worker must defer wg.Done(), perform
	//       increments calls to counter.Add(1), then wait and return Load().
	wg.Wait()
	return counter.Load()
}
