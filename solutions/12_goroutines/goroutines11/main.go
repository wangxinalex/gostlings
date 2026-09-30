// Concept: splitting a slice across per-goroutine partial sums
// Task: sum every value with at most workers goroutines and add the partial totals after joining

package main

import "sync"

func sumParallel(values []int, workers int) int {
	if workers < 1 {
		workers = 1
	}
	if workers > len(values) {
		workers = len(values)
	}

	partials := make([]int, workers)
	var wg sync.WaitGroup
	for slot := range partials {
		wg.Add(1)
		go func(slot int) {
			defer wg.Done()
			for index := slot; index < len(values); index += workers {
				partials[slot] += values[index]
			}
		}(slot)
	}
	wg.Wait()

	total := 0
	for _, partial := range partials {
		total += partial
	}
	return total
}

func main() {}
