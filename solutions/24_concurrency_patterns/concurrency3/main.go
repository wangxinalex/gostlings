// Concept: atomic counters combine lock-free increments with an explicit join.
// Task: count every increment from every worker and return only after all workers join.
// Hint: use atomic.Int64.Add in each worker and Wait before Load.
package main

import (
	"sync"
	"sync/atomic"
)

func incrementConcurrently(workers, increments int) int64 {
	var counter atomic.Int64
	var wg sync.WaitGroup
	wg.Add(workers)
	for range workers {
		go func() {
			defer wg.Done()
			for range increments {
				counter.Add(1)
			}
		}()
	}
	wg.Wait()
	return counter.Load()
}

func main() {}
