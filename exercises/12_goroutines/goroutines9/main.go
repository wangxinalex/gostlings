// Concept: deferred completion covers early returns
// Task: ensure every launched worker completes its WaitGroup bookkeeping, even when it returns early
// Expected behavior: the stopAt worker leaves its result empty, but every worker is joined before return.
// Hint: put defer wg.Done() first in the worker. The stopAt branch may then return safely.

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
			if job == stopAt {
				return
			}
			results[index] = fmt.Sprintf("job %d done", job)
			// TODO: Defer wg.Done() at the beginning so the early return also completes.
			wg.Done()
		}(index, job)
	}
	wg.Wait()
	return results
}

func main() {}
