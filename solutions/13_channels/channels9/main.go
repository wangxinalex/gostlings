package main

import "fmt"

func complete() <-chan struct{} {
	done := make(chan struct{})
	go func() {
		close(done)
	}()
	return done
}

func main() {
	<-complete()
	fmt.Println("completed")
}
