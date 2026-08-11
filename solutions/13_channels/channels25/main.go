package main

import "fmt"

func merge(inputs ...<-chan int) <-chan int {
	out := make(chan int)
	if len(inputs) == 0 {
		close(out)
		return out
	}

	exited := make(chan struct{}, len(inputs))
	for _, input := range inputs {
		go func(in <-chan int) {
			for value := range in {
				out <- value
			}
			exited <- struct{}{}
		}(input)
	}
	go func() {
		for input := 0; input < len(inputs); input++ {
			<-exited
		}
		close(out)
	}()
	return out
}

func main() {
	in := make(chan int, 2)
	in <- 1
	in <- 2
	close(in)
	for value := range merge(in) {
		fmt.Println(value)
	}
}
