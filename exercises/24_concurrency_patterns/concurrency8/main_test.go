package main

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestRunRateLimitedConsumesTokensAndPreservesResults(t *testing.T) {
	tokens := make(chan struct{}, 3)
	for i := 0; i < 3; i++ {
		tokens <- struct{}{}
	}
	got, err := runRateLimited(context.Background(), tokens, 2, []int{1, 2, 3})
	if err != nil || !reflect.DeepEqual(got, []int{2, 4, 6}) {
		t.Fatalf("runRateLimited() = (%v, %v)", got, err)
	}
}

func TestRunRateLimitedStopsWaitingForTokensAfterCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	result := make(chan error, 1)
	go func() { _, err := runRateLimited(ctx, make(chan struct{}), 1, []int{1}); result <- err }()
	cancel()
	select {
	case err := <-result:
		if err != context.Canceled {
			t.Fatalf("error = %v, want context.Canceled", err)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runRateLimited() remained blocked")
	}
}
