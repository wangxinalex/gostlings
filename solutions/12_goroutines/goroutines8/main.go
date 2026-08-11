// Concept: worker inputs are immutable parameters
// Task: pass each job value into its worker and return one result per job

package main

import (
	"fmt"
	"sync"
)

func runWorkersWithInput(jobs []int) []string {
	results := make([]string, len(jobs))
	var wg sync.WaitGroup
	for index, job := range jobs {
		wg.Add(1)
		go func(index, job int) {
			defer wg.Done()
			results[index] = fmt.Sprintf("job %d received", job)
		}(index, job)
	}
	wg.Wait()
	return results
}

func main() {}
