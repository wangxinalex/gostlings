package main

import (
	"context"
	"testing"
	"time"
)

func TestRunReturnsTimeoutResultAndCallsCancel(t *testing.T) {
	previous := withTimeout
	previousWorkGate := workGate
	workGate = make(chan struct{})
	t.Cleanup(func() { workGate = previousWorkGate })

	created := make(chan struct{}, 1)
	canceled := make(chan struct{}, 1)
	var trigger context.CancelFunc
	withTimeout = func(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
		ctx, cancel := context.WithCancel(parent)
		trigger = cancel
		created <- struct{}{}
		return ctx, func() {
			cancel()
			canceled <- struct{}{}
		}
	}
	t.Cleanup(func() { withTimeout = previous })

	result := make(chan string, 1)
	go func() { result <- run() }()

	select {
	case <-created:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("run() did not create its timeout context")
	}
	select {
	case got := <-result:
		t.Fatalf("run() returned %q before timeout cancellation", got)
	default:
	}

	trigger()
	select {
	case got := <-result:
		if got != "worker: timed out" {
			t.Fatalf("run() = %q, want timeout result", got)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("run() did not return after timeout cancellation")
	}

	select {
	case <-canceled:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("run() did not call its timeout cancel function")
	}
}
