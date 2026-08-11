// Concept: ticker-driven loops must stop and publish completion on every exit.
// Task: consume injected ticks until stop or tick input closes, then close done.
// Hint: select on both channels and defer the ticker cleanup hook.
package main

import "time"

var tickerStopped = func() {}

func runTicker(stop <-chan struct{}, ticks <-chan time.Time) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		defer tickerStopped()
		for {
			select {
			case <-stop:
				return
			case _, ok := <-ticks:
				if !ok {
					return
				}
			}
		}
	}()
	return done
}

func main() {}
