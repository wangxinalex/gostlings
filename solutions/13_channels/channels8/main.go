package main

import (
	"fmt"
	"sort"
	"sync"
)

func run(workers int, jobs []int) []int {
	if workers < 1 {
		workers = 1
	}

	jobsCh := make(chan int)
	results := make(chan int)
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for job := range jobsCh {
				results <- job * job
			}
		}()
	}

	go func() {
		for _, job := range jobs {
			jobsCh <- job
		}
		close(jobsCh)
		wg.Wait()
		close(results)
	}()

	var output []int
	for result := range results {
		output = append(output, result)
	}
	return output
}

func main() {
	results := run(2, []int{1, 2, 3, 4})
	sort.Ints(results)
	fmt.Println(results)
}
