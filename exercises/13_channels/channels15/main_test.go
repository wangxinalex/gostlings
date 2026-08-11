package main

import (
	"testing"
	"time"
)

func TestRunWithDoneSeparatesResultFromCompletion(t *testing.T) {
	done := make(chan struct{})
	result := runWithDone(done)

	select {
	case got := <-result:
		if got != 42 {
			t.Fatalf("runWithDone() result = %d, want 42", got)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runWithDone() did not publish its result")
	}

	select {
	case _, ok := <-result:
		if ok {
			t.Fatal("runWithDone() left an unexpected second result")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runWithDone() did not close its result channel")
	}

	select {
	case <-done:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runWithDone() did not close done after completing")
	}
}
