package main

import (
	"context"
	"testing"
	"time"
)

func TestRunReturnsTimeoutResultAndCallsCancel(t *testing.T) {
	previous := withTimeout
	canceled := make(chan struct{}, 1)
	withTimeout = func(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
		ctx, cancel := context.WithCancel(parent)
		cancel()
		return ctx, func() {
			cancel()
			canceled <- struct{}{}
		}
	}
	t.Cleanup(func() { withTimeout = previous })

	if got := run(); got != "worker: timed out" {
		t.Fatalf("run() = %q, want timeout result", got)
	}
	select {
	case <-canceled:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("run() did not call its timeout cancel function")
	}
}
