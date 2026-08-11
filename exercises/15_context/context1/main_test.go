package main

import (
	"context"
	"testing"
	"time"
)

func TestWorkerStopsWhenContextIsCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if got := worker(ctx); got != "worker: canceled" {
		t.Fatalf("worker() = %q, want cancellation result", got)
	}
}

func TestWorkerDoesNotWaitForBlockedWork(t *testing.T) {
	previous := workGate
	workGate = make(chan struct{})
	t.Cleanup(func() { workGate = previous })

	ctx, cancel := context.WithCancel(context.Background())
	result := make(chan string, 1)
	go func() { result <- worker(ctx) }()
	cancel()

	select {
	case got := <-result:
		if got != "worker: canceled" {
			t.Fatalf("worker() = %q, want cancellation result", got)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("worker() waited for blocked work after cancellation")
	}
}
