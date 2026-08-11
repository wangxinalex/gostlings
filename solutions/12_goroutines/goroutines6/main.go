// Concept: completing each started goroutine exactly once
// Task: visit every job and return the number of workers that completed

package main

import "sync"

func runEach(jobs []int, visit func(int)) int {
	completed := make([]bool, len(jobs))
	var wg sync.WaitGroup
	for index, job := range jobs {
		wg.Add(1)
		go func(index, job int) {
			defer wg.Done()
			visit(job)
			completed[index] = true
		}(index, job)
	}
	wg.Wait()

	count := 0
	for _, done := range completed {
		if done {
			count++
		}
	}
	return count
}

func main() {}
