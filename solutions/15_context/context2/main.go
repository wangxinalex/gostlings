// Concept: context.WithTimeout — automatic cancellation after a deadline
// Task: the worker finishes its work before the parent cancels; wrap the context with a 50ms timeout so it prints "timed out"
// Expected output: worker: timed out
// Hint: context.WithTimeout(ctx, duration) auto-cancels after the duration and the ctx.Done channel fires (Go doc: context)

package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("worker: timed out")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("worker: finished work")
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	go worker(ctx)
	time.Sleep(200 * time.Millisecond)
}
