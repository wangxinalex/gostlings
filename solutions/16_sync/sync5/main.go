// Concept: add all known work before waiting for a worker group.
// Task: apply runTask to every job concurrently and return the sum of its results.
// Hint: use a buffered result channel and a WaitGroup; never call Wait before all Add calls.
package main

import "sync"

var runTask = func(job int) int { return job }

func runTasks(jobs []int) int {
	if len(jobs) == 0 {
		return 0
	}
	results := make(chan int, len(jobs))
	var wg sync.WaitGroup
	wg.Add(len(jobs))
	for _, job := range jobs {
		go func(job int) {
			defer wg.Done()
			results <- runTask(job)
		}(job)
	}
	wg.Wait()
	close(results)
	total := 0
	for result := range results {
		total += result
	}
	return total
}

func main() {}
