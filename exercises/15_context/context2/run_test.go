package main

import (
	"testing"
	"time"
)

func TestRunWaitsForWorkerResult(t *testing.T) {
	start := time.Now()
	if got := run(); got != "worker: timed out" {
		t.Fatalf("run() = %q, want %q", got, "worker: timed out")
	}
	if elapsed := time.Since(start); elapsed > 500*time.Millisecond {
		t.Fatalf("run() took %s; it should wait on the result channel, not a fixed sleep", elapsed)
	}
}
