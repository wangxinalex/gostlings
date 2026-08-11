package main

import (
	"context"
	"testing"
	"time"
)

func TestRunReturnsTimeoutResultAndCallsCancel(t *testing.T) {
	previous := withTimeout
	previousWorker := runWorker
	previousWorkGate := workGate
	workGate = make(chan struct{})
	t.Cleanup(func() { workGate = previousWorkGate })

	created := make(chan struct{}, 1)
	canceled := make(chan struct{}, 1)
	workerCalled := make(chan context.Context, 1)
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
	runWorker = func(ctx context.Context) string {
		workerCalled <- ctx
		<-ctx.Done()
		return "worker: timed out"
	}
	t.Cleanup(func() {
		runWorker = previousWorker
		if trigger != nil {
			trigger()
		}
	})

	result := make(chan string, 1)
	go func() { result <- run() }()

	select {
	case <-created:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("run() did not create its timeout context")
	}
	select {
	case <-workerCalled:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("run() did not start the worker")
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

func TestRunReturnsWorkerResultWhenWorkCompletes(t *testing.T) {
	previous := withTimeout
	previousWorkGate := workGate
	workGate = make(chan struct{})
	close(workGate)
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
	t.Cleanup(func() {
		if trigger != nil {
			trigger()
		}
		withTimeout = previous
	})

	result := make(chan string, 1)
	go func() { result <- run() }()

	select {
	case <-created:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("run() did not create its timeout context")
	}
	select {
	case got := <-result:
		if got != "worker: completed" {
			t.Fatalf("run() = %q, want worker result", got)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("run() did not return the worker result")
	}

	select {
	case <-canceled:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("run() did not call its timeout cancel function")
	}
}
