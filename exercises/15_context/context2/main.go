// Concept: context.WithTimeout — automatic cancellation after a deadline
// Task: wrap the context with a 50ms timeout so the worker returns before its 100ms job finishes
// Expected output: worker: timed out
// Hint: send the worker result through the buffered result channel and receive it; fixed sleeps are not synchronization

package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) string {
	select {
	case <-ctx.Done():
		return "worker: timed out"
	case <-time.After(100 * time.Millisecond):
		return "worker: finished work"
	}
}

func run() string {
	ctx := context.Background()
	// TODO: Replace the line above with context.WithTimeout using 50ms,
	//       and remember to call the returned cancel function.

	result := make(chan string, 1)
	go func() {
		result <- worker(ctx)
	}()
	return <-result
}

func main() {
	fmt.Println(run())
}
