package main

import "fmt"

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for value := range in {
			out <- value * value
		}
	}()
	return out
}

func main() {
	in := make(chan int, 3)
	in <- 1
	in <- 2
	in <- 3
	close(in)
	for value := range square(in) {
		fmt.Println(value)
	}
}
