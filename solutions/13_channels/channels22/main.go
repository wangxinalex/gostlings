package main

import (
	"fmt"
	"sync"
)

func parallel(limit int, jobs []int, work func(int) int) []int {
	if limit < 1 {
		limit = 1
	}

	tokens := make(chan struct{}, limit)
	results := make([]int, len(jobs))
	var wg sync.WaitGroup

	for index, job := range jobs {
		index, job := index, job
		wg.Go(func() {
			tokens <- struct{}{}
			defer func() { <-tokens }()
			results[index] = work(job)
		})
	}
	wg.Wait()
	return results
}

func main() {
	results := parallel(2, []int{1, 2, 3, 4}, func(value int) int { return value * value })
	fmt.Println(results)
}
