package main

import "fmt"

type indexedJob struct {
	index int
	value int
}

type indexedResult struct {
	index int
	value int
}

var processOrderedBounded = func(value int) int { return value * value }

func runOrderedBounded(stop <-chan struct{}, workers int, jobs []int) ([]int, bool) {
	if workers < 1 {
		workers = 1
	}
	jobsCh := make(chan indexedJob)
	results := make(chan indexedResult)
	exited := make(chan struct{}, workers)
	go func() {
		defer close(jobsCh)
		for index, value := range jobs {
			select {
			case <-stop:
				return
			case jobsCh <- indexedJob{index: index, value: value}:
			}
		}
	}()
	for range workers {
		go func() {
			defer func() { exited <- struct{}{} }()
			for {
				select {
				case <-stop:
					return
				default:
				}
				var current indexedJob
				select {
				case <-stop:
					return
				case next, ok := <-jobsCh:
					if !ok {
						return
					}
					current = next
				}
				value := processOrderedBounded(current.value)
				select {
				case <-stop:
					return
				case results <- indexedResult{index: current.index, value: value}:
				}
			}
		}()
	}
	go func() {
		for range workers {
			<-exited
		}
		close(results)
	}()

	ordered := make([]int, len(jobs))
	for result := range results {
		ordered[result.index] = result.value
	}
	select {
	case <-stop:
		return []int{}, false
	default:
		return ordered, true
	}
}

func main() { fmt.Println(runOrderedBounded(make(chan struct{}), 1, nil)) }
