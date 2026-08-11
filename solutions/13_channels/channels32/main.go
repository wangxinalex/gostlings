package main

import "fmt"

type request struct {
	value int
	reply chan int
}

func serve(requests <-chan request) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for request := range requests {
			request.reply <- request.value * 2
		}
	}()
	return done
}

func main() { fmt.Println(serve(nil)) }
