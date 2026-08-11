package main

import "fmt"

var onPoolWorkerStart = func() {}
var onPoolWorkerExit = func() {}

func startPool(workers int) (chan<- int, <-chan int) {
	if workers < 1 {
		workers = 1
	}

	jobs := make(chan int)
	results := make(chan int)
	exited := make(chan struct{}, workers)
	for worker := 0; worker < workers; worker++ {
		go func() {
			onPoolWorkerStart()
			defer func() {
				onPoolWorkerExit()
				exited <- struct{}{}
			}()
			for job := range jobs {
				results <- job * job
			}
		}()
	}
	go func() {
		for worker := 0; worker < workers; worker++ {
			<-exited
		}
		close(results)
	}()
	return jobs, results
}

func main() {
	jobs, results := startPool(2)
	jobs <- 2
	jobs <- 3
	close(jobs)
	for result := range results {
		fmt.Println(result)
	}
}
