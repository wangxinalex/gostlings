package main

import "fmt"

func forward(stop <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			var value int
			select {
			case <-stop:
				return
			case received, ok := <-in:
				if !ok {
					return
				}
				value = received
			}

			select {
			case <-stop:
				return
			case out <- value:
			}
		}
	}()
	return out
}

func main() {
	in := make(chan int, 2)
	in <- 3
	in <- 8
	close(in)
	for value := range forward(make(chan struct{}), in) {
		fmt.Println(value)
	}
}
