package main

import (
	"fmt"
	"sync"
)

func merge(inputs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(inputs))

	for _, input := range inputs {
		go func(in <-chan int) {
			defer wg.Done()
			for value := range in {
				out <- value
			}
		}(input)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	left := make(chan int, 2)
	right := make(chan int, 2)
	left <- 1
	left <- 3
	right <- 2
	right <- 4
	close(left)
	close(right)

	for value := range merge(left, right) {
		fmt.Println(value)
	}
}
