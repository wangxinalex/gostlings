// Concept: a basic sync.WaitGroup lifecycle
// Task: return one completion from every worker before runWorkers returns

package main

import (
	"fmt"
	"sync"
)

func runWorkers(count int) []string {
	results := make([]string, count)
	var wg sync.WaitGroup
	for worker := range results {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			results[index] = fmt.Sprintf("worker %d done", index)
		}(worker)
	}
	wg.Wait()
	return results
}

func main() {}
