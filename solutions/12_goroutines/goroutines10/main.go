// Concept: reviewing a complete goroutine lifecycle
// Task: dynamically launch parameterized workers, defer completion, and join before returning

package main

import (
	"fmt"
	"sync"
)

func reviewRun(jobs []int) []string {
	results := make([]string, len(jobs))
	var wg sync.WaitGroup
	for index, job := range jobs {
		wg.Add(1)
		go func(index, job int) {
			defer wg.Done()
			results[index] = fmt.Sprintf("reviewed %d", job)
		}(index, job)
	}
	wg.Wait()
	return results
}

func main() {}
