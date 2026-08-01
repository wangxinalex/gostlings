// Concept: timeout plus cancellation — stop slow producers after the caller gives up
// Task: add a timeout branch, signal the slow producer to stop, and wait for it before returning
// Expected output: timed out
// Hint: select on time.After(100 * time.Millisecond); close stop exactly once and wait on done so no goroutine is left behind

package main

import (
	"fmt"
	"time"
)

func run() string {
	ch := make(chan string)
	stop := make(chan struct{})
	done := make(chan struct{})

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

	// TODO: Add a 100ms timeout case. On timeout, close stop, wait for done,
	//       and return "timed out". Also clean up the producer if ch wins.
	select {
	case msg := <-ch:
		close(stop)
		<-done
		return msg
	}
}

func main() {
	fmt.Println(run())
}
