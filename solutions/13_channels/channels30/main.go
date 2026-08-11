package main

import "fmt"

var processBoundedJob = func(value int) int { return value * value }
var onBoundedQueue = func(capacity int) {}
var onBoundedProcessStart = func(value int) {}

func runBounded(workers, buffer int, jobs []int) []int {
	if workers < 1 {
		workers = 1
	}
	if buffer < 0 {
		buffer = 0
	}

	jobsCh := make(chan int, buffer)
	onBoundedQueue(cap(jobsCh))
	results := make(chan int)
	exited := make(chan struct{}, workers)
	go func() {
		defer close(jobsCh)
		for _, job := range jobs {
			jobsCh <- job
		}
	}()
	for worker := 0; worker < workers; worker++ {
		go func() {
			defer func() { exited <- struct{}{} }()
			for job := range jobsCh {
				onBoundedProcessStart(job)
				results <- processBoundedJob(job)
			}
		}()
	}
	go func() {
		for worker := 0; worker < workers; worker++ {
			<-exited
		}
		close(results)
	}()

	values := make([]int, 0, len(jobs))
	for result := range results {
		values = append(values, result)
	}
	return values
}

func main() {
	fmt.Println(runBounded(2, 1, []int{1, 2, 3, 4}))
}
