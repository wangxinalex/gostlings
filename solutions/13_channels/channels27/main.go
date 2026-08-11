package main

import "fmt"

type indexedJob struct {
	index int
	value int
}

type indexedResult struct {
	index int
	value int
}

var processOrderedJob = func(value int) int { return value * value }

func runOrdered(workers int, jobs []int) []int {
	if workers < 1 {
		workers = 1
	}

	jobsCh := make(chan indexedJob)
	results := make(chan indexedResult)
	exited := make(chan struct{}, workers)
	go func() {
		defer close(jobsCh)
		for index, value := range jobs {
			jobsCh <- indexedJob{index: index, value: value}
		}
	}()
	for worker := 0; worker < workers; worker++ {
		go func() {
			defer func() { exited <- struct{}{} }()
			for job := range jobsCh {
				results <- indexedResult{index: job.index, value: processOrderedJob(job.value)}
			}
		}()
	}
	go func() {
		for worker := 0; worker < workers; worker++ {
			<-exited
		}
		close(results)
	}()

	ordered := make([]int, len(jobs))
	for result := range results {
		ordered[result.index] = result.value
	}
	return ordered
}

func main() {
	fmt.Println(runOrdered(2, []int{4, 1, 3, 2}))
}
