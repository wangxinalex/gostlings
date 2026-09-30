package main

import "fmt"

type indexedResult struct {
	index int
	value int
}

func parallel(limit int, jobs []int, work func(int) int) []int {
	if len(jobs) == 0 {
		return []int{}
	}
	if limit < 1 {
		limit = 1
	}

	permits := make(chan struct{}, limit)
	for range limit {
		permits <- struct{}{}
	}
	results := make(chan indexedResult)
	exited := make(chan struct{}, len(jobs))
	for index, value := range jobs {
		go func(index, value int) {
			defer func() { exited <- struct{}{} }()
			<-permits
			result := work(value)
			permits <- struct{}{}
			results <- indexedResult{index: index, value: result}
		}(index, value)
	}
	go func() {
		for range jobs {
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

func main() { fmt.Println(parallel(1, nil, func(value int) int { return value })) }
