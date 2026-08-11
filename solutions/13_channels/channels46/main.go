package main

import "fmt"

var onRateLimitBeforeSend = func() {}

func rateLimit(stop <-chan struct{}, tokens <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			var value int
			select {
			case <-stop:
				return
			case next, ok := <-in:
				if !ok {
					return
				}
				value = next
			}
			select {
			case <-stop:
				return
			case <-tokens:
			}
			onRateLimitBeforeSend()
			select {
			case <-stop:
				return
			case out <- value:
			}
		}
	}()
	return out
}

func main() { fmt.Println(rateLimit(make(chan struct{}), nil, nil)) }
