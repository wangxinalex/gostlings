package main

import (
	"context"
	"testing"
	"time"
)

func TestStartIfActiveRejectsCanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	called := make(chan struct{}, 1)
	if startIfActive(ctx, func() { called <- struct{}{} }) {
		t.Fatal("startIfActive() = true for a canceled context")
	}
	select {
	case <-called:
		t.Fatal("work ran for a canceled context")
	default:
	}
}

func TestStartIfActiveRunsWorkForActiveContext(t *testing.T) {
	called := make(chan struct{}, 1)
	if !startIfActive(context.Background(), func() { called <- struct{}{} }) {
		t.Fatal("startIfActive() = false for an active context")
	}
	select {
	case <-called:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("work was not called for an active context")
	}
}
