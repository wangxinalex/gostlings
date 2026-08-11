// Concept: reviewing a complete goroutine lifecycle
// Task: dynamically launch parameterized workers, defer completion, and join before returning
// Expected behavior: every input job produces one reviewed result, including dynamic and empty inputs.
// Hint: pass both the result index and job value to the goroutine. Add before go, defer Done, then Wait after the loop.

package main

import "sync"

func reviewRun(jobs []int) []string {
	results := make([]string, len(jobs))
	var wg sync.WaitGroup
	for index := range jobs {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			// TODO: Pass the job into this worker and record its review result.
			results[index] = ""
		}(index)
	}
	wg.Wait()
	return results
}

func main() {}
