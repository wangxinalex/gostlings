// Concept: a rate limiter still needs cancellation around every potentially blocking channel operation.
// Task: forward token-limited input until input closes or stop closes.
// Expected behavior: stop releases a goroutine waiting for a token and one blocked sending to an abandoned output.
// Hint: select between stop and receiving in; after a value, select between stop and receiving a token; finally
// select between stop and out <- value. The output owner defers close(out).
package main

import "fmt"

var onRateLimitBeforeSend = func() {}

func rateLimit(stop <-chan struct{}, tokens <-chan struct{}, in <-chan int) <-chan int {
	return nil // TODO: make input, token, and output waits stop-aware
}

func main() { fmt.Println(rateLimit(make(chan struct{}), nil, nil)) }
