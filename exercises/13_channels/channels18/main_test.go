package main

import (
	"testing"
	"time"
)

func TestRunStopsWorkersAfterFirstError(t *testing.T) {
	jobs := []job{
		{value: 1},
		{fail: true},
		{value: 2},
		{value: 3},
	}
	result := make(chan error, 1)
	go func() { result <- run(3, jobs) }()

	select {
	case err := <-result:
		if err == nil {
			t.Fatal("run() returned nil error for a failing job")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("run() did not finish after cancellation")
	}
}

func TestRunReturnsNilWithoutErrors(t *testing.T) {
	if err := run(2, []job{{value: 1}, {value: 2}}); err != nil {
		t.Fatalf("run() = %v, want nil", err)
	}
}
