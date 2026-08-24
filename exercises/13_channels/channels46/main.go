// Advanced reinforcement: every wait in a token-based rate limiter needs an
// escape path.
// Problem: cancellation may arrive while waiting for input, a rate token, or a
// downstream receiver.
// Without this pattern: the limiter leaks at whichever channel operation is left
// as a plain blocking receive or send.
// Channels: stop cancels; in carries data; tokens carries permits; out carries
// forwarded data and is closed by the output owner.
// Timeline: select in/stop -> select token/stop -> select out/stop -> repeat/close(out)
// Hint: protect input, token, and output in separate selects. After a value,
// select between stop and receiving a token; then select between stop and
// `out <- value`. Defer close(out) in the owner goroutine.
package main

import "fmt"

var onRateLimitBeforeSend = func() {}

func rateLimit(stop <-chan struct{}, tokens <-chan struct{}, in <-chan int) <-chan int {
	return nil // TODO: make input, token, and output waits stop-aware
}

func main() { fmt.Println(rateLimit(make(chan struct{}), nil, nil)) }
