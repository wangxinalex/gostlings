package main

import "fmt"

func drainClosed(ch <-chan int) []int {
	var values []int
	for {
		value, ok := <-ch
		if !ok {
			return values
		}
		values = append(values, value)
	}
}

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	fmt.Println(drainClosed(ch))
}
