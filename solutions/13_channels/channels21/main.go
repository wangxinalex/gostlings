package main

import "fmt"

var onSquareWorkerStart = func() {}
var onSquareWorkerExit = func() {}

func squareWorkers(workers int, jobs <-chan int) <-chan int {
	if workers < 1 {
		workers = 1
	}

	out := make(chan int)
	exited := make(chan struct{}, workers)
	for worker := 0; worker < workers; worker++ {
		go func() {
			onSquareWorkerStart()
			defer func() {
				onSquareWorkerExit()
				exited <- struct{}{}
			}()
			for job := range jobs {
				out <- job * job
			}
		}()
	}
	go func() {
		for worker := 0; worker < workers; worker++ {
			<-exited
		}
		close(out)
	}()
	return out
}

func main() {
	jobs := make(chan int, 2)
	jobs <- 2
	jobs <- 3
	close(jobs)
	for result := range squareWorkers(2, jobs) {
		fmt.Println(result)
	}
}
