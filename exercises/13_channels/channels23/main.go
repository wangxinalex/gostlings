// Concept: ticker channels can act as a rate limiter
// Task: forward one input only after one tick is received
// Expected behavior: every input consumes one tick, then output closes
// Hint: wait on <-ticks before sending each value; the caller owns the ticker lifecycle

package main

import (
	"fmt"
	"time"
)

func rateLimit(ticks <-chan time.Time, in <-chan int) <-chan int {
	// Thought: model “permission to start the next job” as a channel event, so
	// production follows the ticker cadence instead of a fixed Sleep.
	return nil // TODO: wait for a tick before forwarding each input
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
