// Concept: context.WithTimeout limits a request and releases its resources.
// Task: create a timeout context, defer its cancel function, and return the
// worker result through a capacity-one channel.
// Hint: context.WithTimeout returns (context.Context, context.CancelFunc).
// A buffered result channel lets the worker report once even as cancellation wins.
package main

import (
	"context"
	"fmt"
	"time"
)

var withTimeout = context.WithTimeout
var workGate = make(chan struct{})

func worker(ctx context.Context) string {
	select {
	case <-ctx.Done():
		return "worker: timed out"
	case <-workGate:
		return "worker: completed"
	}
}

func run() string {
	// TODO: Use withTimeout, defer cancel, and receive from a capacity-one result channel.
	timeout := 50 * time.Millisecond
	_ = timeout
	return "worker: completed"
}

func main() { fmt.Println(run()) }
