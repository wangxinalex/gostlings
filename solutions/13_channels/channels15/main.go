package main

import (
	"fmt"
	"sync"
)

func merge(stop <-chan struct{}, inputs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, input := range inputs {
		in := input
		wg.Go(func() {
			for {
				select {
				case <-stop:
					return
				case value, ok := <-in:
					if !ok {
						return
					}
					select {
					case out <- value:
					case <-stop:
						return
					}
				}
			}
		})
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	input := make(chan int, 2)
	input <- 1
	input <- 2
	close(input)
	for value := range merge(make(chan struct{}), input) {
		fmt.Println(value)
	}
}
