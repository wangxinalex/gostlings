// Concept: context.WithCancel for cancellation signals
// Task: call cancel when the work is done so the goroutine receives the cancellation signal
// Expected output: worker: received cancel signal
// Hint: context.WithCancel returns a context and a cancel function; call cancel when you want the goroutine to stop (Go doc: context)

package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("worker: received cancel signal")
	case <-time.After(2 * time.Second):
		fmt.Println("worker: work completed")
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(100 * time.Millisecond)
	cancel()
	time.Sleep(100 * time.Millisecond)
}
