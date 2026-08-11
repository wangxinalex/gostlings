package main

import (
	"context"
	"testing"
	"time"
)

func TestRunUntilUsesAbsoluteDeadlineAndCallsCancel(t *testing.T) {
	previous := withDeadline
	previousWorkGate := workGate
	workGate = make(chan struct{})
	t.Cleanup(func() { workGate = previousWorkGate })

	created := make(chan struct{}, 1)
	called := make(chan struct{}, 1)
	var gotParent context.Context
	var gotDeadline time.Time
	var trigger context.CancelFunc
	withDeadline = func(parent context.Context, deadline time.Time) (context.Context, context.CancelFunc) {
		gotParent = parent
		gotDeadline = deadline
		ctx, cancel := context.WithCancel(parent)
		trigger = cancel
		created <- struct{}{}
		return ctx, func() {
			cancel()
			called <- struct{}{}
		}
	}
	t.Cleanup(func() { withDeadline = previous })

	parent := context.Background()
	deadline := time.Now().Add(time.Hour)
	result := make(chan string, 1)
	go func() { result <- runUntil(parent, deadline) }()

	select {
	case <-created:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runUntil() did not create its deadline context")
	}
	if gotParent != parent || !gotDeadline.Equal(deadline) {
		t.Fatalf("withDeadline() = (%v, %v), want supplied parent and deadline", gotParent, gotDeadline)
	}
	select {
	case got := <-result:
		t.Fatalf("runUntil() returned %q before deadline cancellation", got)
	default:
	}

	trigger()
	select {
	case got := <-result:
		if got != "work: deadline exceeded" {
			t.Fatalf("runUntil() = %q, want deadline result", got)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runUntil() did not return after deadline cancellation")
	}

	select {
	case <-called:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runUntil() did not call its deadline cancel function")
	}
}

func TestRunUntilPreservesParentCancellation(t *testing.T) {
	parent, cancel := context.WithCancel(context.Background())
	cancel()
	result := make(chan string, 1)
	go func() { result <- runUntil(parent, time.Now().Add(time.Hour)) }()
	select {
	case got := <-result:
		if got != "work: deadline exceeded" {
			t.Fatalf("runUntil() = %q, want cancellation result from the parent", got)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runUntil() did not preserve parent cancellation")
	}
}
