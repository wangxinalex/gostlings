// Concept: waiting for a dynamic number of jobs
// Task: run work once per job and return only after every job completes

package main

import "sync"

func runJobs(jobs []int, work func(int) string) []string {
	results := make([]string, len(jobs))
	var wg sync.WaitGroup
	for index, job := range jobs {
		wg.Add(1)
		go func(index, job int) {
			defer wg.Done()
			results[index] = work(job)
		}(index, job)
	}
	wg.Wait()
	return results
}

func main() {}
