package main

import (
	"fmt"
	"time"
)

func run(done chan struct{}) string {
	ch := make(chan string)
	stop := make(chan struct{})

	go func() {
		defer close(done)
		timer := time.NewTimer(2 * time.Second)
		defer timer.Stop()

		select {
		case <-timer.C:
		case <-stop:
			return
		}

		select {
		case ch <- "late result":
		case <-stop:
		}
	}()

	select {
	case value := <-ch:
		close(stop)
		<-done
		return value
	case <-time.After(50 * time.Millisecond):
		close(stop)
		<-done
		return "timed out"
	}
}

func main() {
	done := make(chan struct{})
	fmt.Println(run(done))
}
