// Concept: broadcast cancellation stops a group of goroutines
// Task: start count workers and close the returned done channel after all stop
// Expected behavior: closing stop lets every worker exit and done closes
// Hint: each worker waits on stop; a WaitGroup coordinator closes done once

package main

import (
	"fmt"
	"time"
)

func startWorkers(count int, stop <-chan struct{}) <-chan struct{} {
	done := make(chan struct{})
	// Thought: stop is a broadcast signal; every worker observes the same close
	// event. done indicates that the whole worker group has finished.
	return done // TODO: start workers and close done after all of them exit
}

func main() {
	stop := make(chan struct{})
	done := startWorkers(3, stop)
	close(stop)
	select {
	case <-done:
		fmt.Println("workers stopped")
	case <-time.After(time.Second):
		fmt.Println("workers still running")
	}
}
