// Concept: ticker-driven loops must stop and publish completion on every exit.
// Task: consume injected ticks until stop or tick input closes, then close done.
// Hint: select on both channels and defer the ticker cleanup hook.
package main

import "time"

var tickerStopped = func() {}

func runTicker(stop <-chan struct{}, ticks <-chan time.Time) <-chan struct{} {
	// TODO: Stop consuming on either exit signal and close done exactly once.
	return nil
}

func main() {}
