package main

import "fmt"

var onMergeBeforeSend = func() {}

func merge(stop <-chan struct{}, inputs ...<-chan int) <-chan int {
	out := make(chan int)
	if len(inputs) == 0 {
		close(out)
		return out
	}

	exited := make(chan struct{}, len(inputs))
	for _, input := range inputs {
		go func(in <-chan int) {
			defer func() { exited <- struct{}{} }()
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

				onMergeBeforeSend()
				select {
				case <-stop:
					return
				case out <- value:
				}
			}
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
	for value := range merge(make(chan struct{}), in) {
		fmt.Println(value)
	}
}
