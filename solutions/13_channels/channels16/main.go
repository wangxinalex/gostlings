package main

import "fmt"

func produce(stop <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for value := 1; ; value++ {
			select {
			case <-stop:
				return
			default:
			}

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
	stop := make(chan struct{})
	out := produce(stop)
	fmt.Println(<-out)
	close(stop)
}
