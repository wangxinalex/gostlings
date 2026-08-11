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
var runWorker = worker

func worker(ctx context.Context) string {
	select {
	case <-ctx.Done():
		return "worker: timed out"
	case <-workGate:
		return "worker: completed"
	}
}

func run() string {
	ctx, cancel := withTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	result := make(chan string, 1)
	go func() { result <- runWorker(ctx) }()
	return <-result
}

func main() { fmt.Println(run()) }
