package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRunRequestCompletesAndJoinsWorkers(t *testing.T) {
	const workers = 3
	previousTimeout, previousWork := withRequestTimeout, requestWork
	started := make(chan struct{}, workers)
	stopped := make(chan struct{}, workers)
	canceled := make(chan struct{}, 1)
	withRequestTimeout = func(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
		if timeout <= 0 {
			t.Errorf("timeout = %v, want positive timeout", timeout)
		}
		ctx, cancel := context.WithCancel(parent)
		return ctx, func() { cancel(); canceled <- struct{}{} }
	}
	requestWork = func(ctx context.Context, id int) (string, error) {
		started <- struct{}{}
		stopped <- struct{}{}
		return "ok", nil
	}
	t.Cleanup(func() { withRequestTimeout, requestWork = previousTimeout, previousWork })

	got, err := runRequest(context.Background(), workers)
	if err != nil || got != "processed 3 requests" {
		t.Fatalf("runRequest() = (%q, %v), want successful summary", got, err)
	}
	for i := 0; i < workers; i++ {
		select {
		case <-started:
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("worker %d did not start", i)
		}
		select {
		case <-stopped:
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("worker %d did not stop", i)
		}
	}
	select {
	case <-canceled:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runRequest() did not cancel its derived context")
	}
}

func TestRunRequestCancelsBlockedWorkersBeforeReturning(t *testing.T) {
	const workers = 3
	previousTimeout, previousWork := withRequestTimeout, requestWork
	started := make(chan struct{}, workers)
	stopped := make(chan struct{}, workers)
	requestWork = func(ctx context.Context, id int) (string, error) {
		started <- struct{}{}
		<-ctx.Done()
		stopped <- struct{}{}
		return "", ctx.Err()
	}
	withRequestTimeout = func(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
		return context.WithCancel(parent)
	}
	t.Cleanup(func() { withRequestTimeout, requestWork = previousTimeout, previousWork })

	parent, cancel := context.WithCancel(context.Background())
	result := make(chan error, 1)
	go func() {
		_, err := runRequest(parent, workers)
		result <- err
	}()
	for i := 0; i < workers; i++ {
		select {
		case <-started:
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("worker %d did not start", i)
		}
	}
	cancel()
	select {
	case err := <-result:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("runRequest() error = %v, want context.Canceled", err)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runRequest() did not return after cancellation")
	}
	for i := 0; i < workers; i++ {
		select {
		case <-stopped:
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("worker %d was not joined after cancellation", i)
		}
	}
}

func TestRunRequestRejectsAnInvalidWorkerCount(t *testing.T) {
	if _, err := runRequest(context.Background(), 0); err == nil {
		t.Fatal("runRequest() with zero workers returned nil error")
	}
}
