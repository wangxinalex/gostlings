package main

import "fmt"

func relay(stop <-chan struct{}, in <-chan int) <-chan int {
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
			default:
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
func main() { fmt.Println(relay(make(chan struct{}), nil)) }
