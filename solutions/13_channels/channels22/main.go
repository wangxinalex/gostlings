package main

import "fmt"

var onForwarderExit = func() {}

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
			onForwarderExit()
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
	first := make(chan int, 1)
	second := make(chan int, 1)
	first <- 1
	second <- 2
	close(first)
	close(second)
	for value := range merge(first, second) {
		fmt.Println(value)
	}
}
