package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestShutdownReturnsSuccessfulResults(t *testing.T) {
	got, err := shutdown(context.Background(), 2, []shutdownJob{{value: 1}, {value: 3}})
	if err != nil || len(got) != 2 || got[0].value != 2 || got[1].value != 6 {
		t.Fatalf("shutdown() = (%v, %v)", got, err)
	}
}

func TestShutdownReturnsCancellationAfterWorkersStop(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	result := make(chan error, 1)
	go func() { _, err := shutdown(ctx, 2, []shutdownJob{{value: 1}, {value: 2}}); result <- err }()
	cancel()
	select {
	case err := <-result:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v, want cancellation", err)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("shutdown() did not join after cancellation")
	}
}
