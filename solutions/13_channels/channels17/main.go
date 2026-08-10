package main

import (
	"fmt"
	"sync"
)

type indexedJob struct {
	index int
	value int
}

type indexedResult struct {
	index int
	value int
}

func runOrdered(workers int, jobs []int) []int {
	if workers < 1 {
		workers = 1
	}

	jobsCh := make(chan indexedJob)
	results := make(chan indexedResult)
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Go(func() {
			for job := range jobsCh {
				results <- indexedResult{index: job.index, value: job.value * job.value}
			}
		})
	}

	go func() {
		for index, value := range jobs {
			jobsCh <- indexedJob{index: index, value: value}
		}
		close(jobsCh)
		wg.Wait()
		close(results)
	}()

	output := make([]int, len(jobs))
	for result := range results {
		output[result.index] = result.value
	}
	return output
}

func main() {
	fmt.Println(runOrdered(2, []int{1, 2, 3, 4}))
}
