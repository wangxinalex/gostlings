package main

import "fmt"

func complete() <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
	}()
	return done
}

func main() {
	<-complete()
	fmt.Println("complete")
}
