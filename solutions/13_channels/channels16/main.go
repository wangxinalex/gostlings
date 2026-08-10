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

	for i := 0; i < workers; i++ {
		wg.Go(func() {
			for job := range jobsCh {
				results <- job * job
			}
		})
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
	output := run(2, []int{1, 2, 3, 4})
	sort.Ints(output)
	fmt.Println(output)
}
