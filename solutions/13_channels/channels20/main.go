package main

import "fmt"

func transform(in <-chan int, fn func(int) int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for value := range in {
			out <- fn(value)
		}
	}()
	return out
}

func pipeline(in <-chan int) <-chan int {
	doubled := transform(in, func(value int) int { return value * 2 })
	return transform(doubled, func(value int) int { return value + 1 })
}

func main() {
	in := make(chan int, 3)
	in <- 1
	in <- 2
	in <- 3
	close(in)
	for value := range pipeline(in) {
		fmt.Println(value)
	}
}
