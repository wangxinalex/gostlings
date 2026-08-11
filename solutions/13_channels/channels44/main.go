package main

import "fmt"

type indexedResult struct {
	index int
	value int
}

func parallel(stop <-chan struct{}, limit int, jobs []int, work func(int) int) ([]int, bool) {
	if len(jobs) == 0 {
		select {
		case <-stop:
			return []int{}, false
		default:
			return []int{}, true
		}
	}
	if limit < 1 {
		limit = 1
	}

	tokens := make(chan struct{}, limit)
	for range limit {
		tokens <- struct{}{}
	}
	results := make(chan indexedResult)
	exited := make(chan struct{}, len(jobs))
	for index, value := range jobs {
		go func(index, value int) {
			defer func() { exited <- struct{}{} }()
			select {
			case <-stop:
				return
			case <-tokens:
			}
			result := work(value)
			tokens <- struct{}{}
			select {
			case <-stop:
				return
			case results <- indexedResult{index: index, value: result}:
			}
		}(index, value)
	}
	go func() {
		for range jobs {
			<-exited
		}
		close(results)
	}()

	ordered := make([]int, len(jobs))
	count := 0
	for result := range results {
		ordered[result.index] = result.value
		count++
	}
	select {
	case <-stop:
		return ordered[:count], false
	default:
		return ordered, count == len(jobs)
	}
}

func main() { fmt.Println(parallel(make(chan struct{}), 1, nil, func(value int) int { return value })) }
