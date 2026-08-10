package main

import (
	"fmt"
	"sort"
	"sync"
)

func squareWorkers(workers int, jobs <-chan int) <-chan int {
	if workers < 1 {
		workers = 1
	}
	results := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Go(func() {
			for job := range jobs {
				results <- job * job
			}
		})
	}

	go func() {
		wg.Wait()
		close(results)
	}()
	return results
}

func main() {
	jobs := make(chan int, 3)
	jobs <- 1
	jobs <- 2
	jobs <- 3
	close(jobs)

	var results []int
	for result := range squareWorkers(2, jobs) {
		results = append(results, result)
	}
	sort.Ints(results)
	fmt.Println(results)
}
