package main

import (
	"testing"
	"time"
)

func TestRunTimesOutWithoutWaitingForSlowProducer(t *testing.T) {
	start := time.Now()
	if got := run(); got != "timed out" {
		t.Fatalf("run() = %q, want %q", got, "timed out")
	}
	if elapsed := time.Since(start); elapsed > 500*time.Millisecond {
		t.Fatalf("run() took %s; timeout should stop the slow producer", elapsed)
	}
}
