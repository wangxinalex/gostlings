package main

import "fmt"

func produce(stop <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, value := range []int{1, 2, 3} {
			select {
			case out <- value:
			case <-stop:
				return
			}
		}
	}()
	return out
}

func main() { fmt.Println(produce(make(chan struct{}))) }
