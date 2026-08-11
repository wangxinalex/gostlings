package main

import (
	"testing"
	"time"
)

func TestAwaitReturnsAReadyResult(t *testing.T) {
	result := make(chan string, 1)
	result <- "ready"

	if got := await(result); got != "ready" {
		t.Fatalf("await() = %q, want %q", got, "ready")
	}
}

func TestAwaitTimesOutForASilentChannel(t *testing.T) {
	returned := make(chan string, 1)
	go func() { returned <- await(make(chan string)) }()

	select {
	case got := <-returned:
		if got != "timed out" {
			t.Fatalf("await() = %q, want %q", got, "timed out")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("await() did not time out for a silent channel")
	}
}
