package main

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestRunPoolReturnsEverySortedResult(t *testing.T) {
	if got, err := runPool(context.Background(), 4, 2, []int{1, 3, 2}); err != nil || !reflect.DeepEqual(got, []int{2, 4, 6}) {
		t.Fatalf("runPool() = (%v, %v), want [2 4 6], nil", got, err)
	}
}

func TestRunPoolJoinsBlockedWorkAfterCancellation(t *testing.T) {
	previous := poolWork
	started := make(chan struct{}, 2)
	poolWork = func(ctx context.Context, job int) (int, error) {
		started <- struct{}{}
		<-ctx.Done()
		return 0, ctx.Err()
	}
	t.Cleanup(func() { poolWork = previous })
	ctx, cancel := context.WithCancel(context.Background())
	result := make(chan error, 1)
	go func() { _, err := runPool(ctx, 2, 2, []int{1, 2}); result <- err }()
	for i := 0; i < 2; i++ {
		select {
		case <-started:
		case <-time.After(500 * time.Millisecond):
			t.Fatal("worker did not start")
		}
	}
	cancel()
	select {
	case err := <-result:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("runPool() error = %v, want context.Canceled", err)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runPool() did not join after cancellation")
	}
}
