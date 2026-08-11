package main

import "fmt"

func run(workers int, jobs []int) []int {
	if workers < 1 || len(jobs) == 0 {
		return []int{}
	}
	jobsCh, results := make(chan int), make(chan int)
	exited := make(chan struct{}, workers)
	go func() {
		defer close(jobsCh)
		for _, job := range jobs {
			jobsCh <- job
		}
	}()
	for range workers {
		go func() {
			defer func() { exited <- struct{}{} }()
			for job := range jobsCh {
				results <- job * job
			}
		}()
	}
	go func() {
		for range workers {
			<-exited
		}
		close(results)
	}()
	values := make([]int, 0, len(jobs))
	for value := range results {
		values = append(values, value)
	}
	return values
}

func main() { fmt.Println(run(2, []int{1, 2, 3})) }
