package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestReceiveReturnsCancellationWhileInputIsBlocked(t *testing.T) {
	in := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	result := make(chan error, 1)
	go func() {
		_, err := receive(ctx, in)
		result <- err
	}()
	cancel()
	select {
	case err := <-result:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("receive() error = %v, want context.Canceled", err)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("receive() remained blocked after cancellation")
	}
}

func TestReceiveReturnsAnInputValue(t *testing.T) {
	in := make(chan int, 1)
	in <- 42
	value, err := receive(context.Background(), in)
	if err != nil || value != 42 {
		t.Fatalf("receive() = (%d, %v), want (42, nil)", value, err)
	}
}
