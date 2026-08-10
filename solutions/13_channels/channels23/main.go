package main

import (
	"fmt"
	"time"
)

func rateLimit(ticks <-chan time.Time, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for value := range in {
			<-ticks
			out <- value
		}
	}()
	return out
}

func main() {
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()
	in := make(chan int, 2)
	in <- 1
	in <- 2
	close(in)
	for value := range rateLimit(ticker.C, in) {
		fmt.Println(value)
	}
}
