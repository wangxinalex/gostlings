package main

import (
	"fmt"
	"sync"
)

func merge(inputs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, input := range inputs {
		in := input
		wg.Go(func() {
			for value := range in {
				out <- value
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
