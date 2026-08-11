package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRetryRetriesThenSucceeds(t *testing.T) {
	attempts := 0
	backoff := make(chan time.Duration, 1)
	backoff <- 0
	err := retry(context.Background(), 2, backoff, func() error {
		attempts++
		if attempts == 1 {
			return retryableError{}
		}
		return nil
	})
	if err != nil || attempts != 2 {
		t.Fatalf("retry() = %v after %d attempts", err, attempts)
	}
}

func TestRetryStopsBackoffOnCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	backoff := make(chan time.Duration)
	result := make(chan error, 1)
	go func() { result <- retry(ctx, 2, backoff, func() error { return retryableError{} }) }()
	cancel()
	select {
	case err := <-result:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v, want cancellation", err)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("retry() remained blocked on backoff")
	}
}

func TestRetryReturnsPermanentErrorImmediately(t *testing.T) {
	want := errors.New("permanent")
	if err := retry(context.Background(), 3, make(chan time.Duration), func() error { return want }); !errors.Is(err, want) {
		t.Fatalf("error = %v, want permanent error", err)
	}
}
