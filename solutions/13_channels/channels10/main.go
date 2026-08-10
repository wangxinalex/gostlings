package main

import "fmt"

func produce(stop <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for value := 1; value <= 3; value++ {
			select {
			case out <- value:
			case <-stop:
				return
			}
		}
	}()
	return out
}

func main() {
	for value := range produce(make(chan struct{})) {
		fmt.Println(value)
	}
}
