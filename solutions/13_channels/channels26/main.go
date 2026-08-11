package main

import "fmt"

var processJob = func(value int) int { return value * value }

func run(workers int, jobs []int) []int {
	if workers < 1 {
		workers = 1
	}

	jobsCh := make(chan int)
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
				results <- processJob(job)
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
	fmt.Println(run(2, []int{1, 2, 3, 4}))
}
