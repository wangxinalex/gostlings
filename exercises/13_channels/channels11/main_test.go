package main

import (
	"testing"
	"time"
)

func TestRunTimesOutAndReturnsPromptly(t *testing.T) {
	start := time.Now()
	done := make(chan struct{})
	if got := run(done); got != "timed out" {
		t.Fatalf("run() = %q, want %q", got, "timed out")
	}
	if elapsed := time.Since(start); elapsed > 500*time.Millisecond {
		t.Fatalf("run() took %s; cancellation did not stop the slow producer", elapsed)
	}
	select {
	case <-done:
	default:
		t.Fatal("run() returned before the producer closed done")
	}
}
