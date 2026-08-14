// Concept: worker inputs are immutable parameters
// Task: pass each job value into its worker and return one result per job
// Expected behavior: every job produces its matching received result.
// Hint: pass job as an argument alongside the result index; do not make the worker look up a changing loop value.
// Go 1.22 gives range variables per-iteration scope, but explicit parameters are
// still the clearest and most portable worker-input pattern.

package main

import "sync"

func runWorkersWithInput(jobs []int) []string {
	results := make([]string, len(jobs))
	var wg sync.WaitGroup
	for index := range jobs {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			// TODO: Accept the job as a worker parameter and format its result.
			results[index] = ""
		}(index)
	}
	wg.Wait()
	return results
}

func main() {}
