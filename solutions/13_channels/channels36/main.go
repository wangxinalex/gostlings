package main

import "fmt"

type result struct {
	value int
	err   error
}

func mergeResults(inputs ...<-chan result) <-chan result {
	out := make(chan result)
	active := 0
	for _, input := range inputs {
		if input != nil {
			active++
		}
	}
	if active == 0 {
		close(out)
		return out
	}
	exited := make(chan struct{}, active)
	for _, input := range inputs {
		if input == nil {
			continue
		}
		go func(input <-chan result) {
			for item := range input {
				out <- item
			}
			exited <- struct{}{}
		}(input)
	}
	go func() {
		for range active {
			<-exited
		}
		close(out)
	}()
	return out
}

func main() { fmt.Println(mergeResults()) }
