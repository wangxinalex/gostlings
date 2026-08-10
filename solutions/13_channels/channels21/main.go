package main

import "fmt"

func transform(stop <-chan struct{}, in <-chan int, fn func(int) int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-stop:
				return
			case value, ok := <-in:
				if !ok {
					return
				}
				select {
				case out <- fn(value):
				case <-stop:
					return
				}
			}
		}
	}()
	return out
}

func pipeline(stop <-chan struct{}, in <-chan int) <-chan int {
	doubled := transform(stop, in, func(value int) int { return value * 2 })
	return transform(stop, doubled, func(value int) int { return value + 1 })
}

func main() {
	in := make(chan int, 2)
	in <- 1
	in <- 2
	close(in)
	for value := range pipeline(make(chan struct{}), in) {
		fmt.Println(value)
	}
}
