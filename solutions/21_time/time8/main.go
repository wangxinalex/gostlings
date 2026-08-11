// Concept: periodic work owns both ticker and deadline cleanup.
// Task: count ticks until context cancellation or the deadline expires.
// Hint: create one ticker and one timer, stop both with defer, and select all exits.
package main

import (
	"context"
	"time"
)

var periodicTick = func() {}

func periodic(ctx context.Context, interval, deadline time.Duration) int {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	timer := time.NewTimer(deadline)
	defer timer.Stop()
	count := 0
	for {
		select {
		case <-ticker.C:
			periodicTick()
			count++
		case <-ctx.Done():
			return count
		case <-timer.C:
			return count
		}
	}
}

func main() {}
