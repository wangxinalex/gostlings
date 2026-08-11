package main

import "fmt"

func runWithDone(done chan struct{}) <-chan int {
	out := make(chan int, 1)
	go func() {
		defer close(done)
		defer close(out)
		out <- 42
	}()
	return out
}

func main() {
	done := make(chan struct{})
	for value := range runWithDone(done) {
		fmt.Println(value)
	}
	<-done
}
