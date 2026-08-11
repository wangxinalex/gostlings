package main

import (
	"context"
	"sync"
	"testing"
)

func TestRunMeasuredCountsCompletedJobs(t *testing.T) {
	previous := measureWork
	measureWork = func(ctx context.Context, job int) bool { return true }
	t.Cleanup(func() { measureWork = previous })
	completed, canceled := runMeasured(context.Background(), 3, []int{1, 2, 3, 4})
	if completed != 4 || canceled != 0 {
		t.Fatalf("runMeasured() = (%d, %d), want (4, 0)", completed, canceled)
	}
}

func TestRunMeasuredCountsCancellationAfterWorkersJoin(t *testing.T) {
	previous := measureWork
	started := make(chan struct{}, 2)
	measureWork = func(ctx context.Context, job int) bool { started <- struct{}{}; <-ctx.Done(); return false }
	t.Cleanup(func() { measureWork = previous })
	ctx, cancel := context.WithCancel(context.Background())
	result := make(chan [2]int64, 1)
	go func() {
		completed, canceled := runMeasured(ctx, 2, []int{1, 2})
		result <- [2]int64{completed, canceled}
	}()
	for i := 0; i < 2; i++ {
		<-started
	}
	cancel()
	got := <-result
	if got != [2]int64{0, 2} {
		t.Fatalf("runMeasured() = %v, want [0 2]", got)
	}
}

func TestRunMeasuredHandlesEmptyJobs(t *testing.T) {
	var wg sync.WaitGroup
	_ = wg
	completed, canceled := runMeasured(context.Background(), 1, nil)
	if completed != 0 || canceled != 0 {
		t.Fatalf("empty run = (%d, %d)", completed, canceled)
	}
}
