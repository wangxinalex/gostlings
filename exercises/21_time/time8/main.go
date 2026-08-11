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
	// TODO: Count periodic ticks until ctx or the deadline ends, stopping resources on return.
	return 0
}

func main() {}
