package main

import (
	"testing"
	"time"
)

func TestAwaitReturnsValueWhenReady(t *testing.T) {
	ch := make(chan string, 1)
	ch <- "ready"

	if got := await(ch); got != "ready" {
		t.Fatalf("await() = %q, want %q", got, "ready")
	}
}

func TestAwaitTimesOutWhenInputStaysSilent(t *testing.T) {
	start := time.Now()
	if got := await(make(chan string)); got != "timed out" {
		t.Fatalf("await() = %q, want %q", got, "timed out")
	}
	if elapsed := time.Since(start); elapsed > 300*time.Millisecond {
		t.Fatalf("await() took %s; want a short timeout", elapsed)
	}
}
