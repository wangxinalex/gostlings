// Concept: deferred completion covers early returns
// Task: ensure every launched worker completes its WaitGroup bookkeeping, even when it returns early

package main

import (
	"fmt"
	"sync"
)

func runWithEarlyReturn(jobs []int, stopAt int) []string {
	results := make([]string, len(jobs))
	var wg sync.WaitGroup
	for index, job := range jobs {
		wg.Add(1)
		go func(index, job int) {
			defer wg.Done()
			if job == stopAt {
				return
			}
			results[index] = fmt.Sprintf("job %d done", job)
		}(index, job)
	}
	wg.Wait()
	return results
}

func main() {}
