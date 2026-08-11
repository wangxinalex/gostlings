package main

import "fmt"

func collectOrStop(stop <-chan struct{}, work <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			var value int
			select {
			case <-stop:
				return
			case received, ok := <-work:
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
func main() { fmt.Println(collectOrStop(make(chan struct{}), nil)) }
