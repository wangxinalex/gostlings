package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRunServiceReturnsOrderedResponses(t *testing.T) {
	got, err := runService(context.Background(), 3, 2, []request{{value: 2}, {value: 1}, {value: 3}})
	if err != nil || len(got) != 3 || got[0].value != 4 || got[1].value != 2 || got[2].value != 6 {
		t.Fatalf("runService() = (%v, %v)", got, err)
	}
}

func TestRunServiceStopsOnFirstFailureAndCancellation(t *testing.T) {
	_, err := runService(context.Background(), 2, 1, []request{{value: 1}, {fail: true}, {value: 3}})
	if !errors.Is(err, errRequestFailed) {
		t.Fatalf("error = %v, want request failure", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	result := make(chan error, 1)
	go func() { _, err := runService(ctx, 2, 1, []request{{value: 1}}); result <- err }()
	select {
	case err := <-result:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("canceled error = %v", err)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runService() did not stop")
	}
}
