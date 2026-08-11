package main

import (
	"context"
	"testing"
	"time"
)

func TestPeriodicStopsOnContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	previous := periodicTick
	seen := make(chan struct{}, 1)
	periodicTick = func() { seen <- struct{}{} }
	t.Cleanup(func() { periodicTick = previous })
	result := make(chan int, 1)
	go func() { result <- periodic(ctx, time.Millisecond, time.Second) }()
	select {
	case <-seen:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("periodic() did not produce a tick")
	}
	cancel()
	select {
	case count := <-result:
		if count < 0 {
			t.Fatalf("periodic() count = %d, want non-negative", count)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("periodic() did not stop after cancellation")
	}
}

func TestPeriodicStopsAtDeadline(t *testing.T) {
	start := time.Now()
	count := periodic(context.Background(), time.Millisecond, 5*time.Millisecond)
	if count < 0 || time.Since(start) > 500*time.Millisecond {
		t.Fatalf("periodic() count=%d did not honor deadline promptly", count)
	}
}
