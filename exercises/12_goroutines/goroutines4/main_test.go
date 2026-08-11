package main

import "testing"

func TestRunWorkersWaitsForEveryWorker(t *testing.T) {
	got := runWorkers(4)
	if len(got) != 4 {
		t.Fatalf("runWorkers(4) returned %d completions, want 4", len(got))
	}
}

func TestRunWorkersHandlesZeroWorkers(t *testing.T) {
	if got := runWorkers(0); len(got) != 0 {
		t.Fatalf("runWorkers(0) returned %d completions, want 0", len(got))
	}
}
