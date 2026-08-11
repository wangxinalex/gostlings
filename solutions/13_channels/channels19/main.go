package main

import (
	"fmt"
	"time"
)

func run(done chan struct{}) string {
	result := make(chan string)
	stop := make(chan struct{})
	go func() {
		defer close(done)
		select {
		case <-time.After(100 * time.Millisecond):
		case <-stop:
			return
		}

		select {
		case result <- "finished":
		case <-stop:
		}
	}()

	select {
	case value := <-result:
		close(stop)
		<-done
		return value
	case <-time.After(25 * time.Millisecond):
		close(stop)
		<-done
		return "timed out"
	}
}

func main() {
	done := make(chan struct{})
	fmt.Println(run(done))
}
