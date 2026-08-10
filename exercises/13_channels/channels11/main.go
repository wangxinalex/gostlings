// Concept: timeout plus cancellation plus join
// Task: on timeout, cancel the slow producer, wait for done, and return "timed out"
// Expected behavior: run returns quickly and done is closed before it returns
// Hint: select on the result and time.After; close stop exactly once, then receive done

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

	// TODO: Add a 50ms timeout branch. On timeout, close stop, wait for done,
	// and return "timed out". Also clean up the producer if ch wins.
	return <-ch
}

func main() {
	done := make(chan struct{})
	fmt.Println(run(done))
}
