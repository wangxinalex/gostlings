package main

import "testing"

func TestIncrementConcurrentlyReturnsExactCount(t *testing.T) {
	const (
		workers    = 100
		increments = 1000
	)
	const want = int64(workers * increments)

	for attempt := 0; attempt < 5; attempt++ {
		if got := incrementConcurrently(workers, increments); got != want {
			t.Fatalf("attempt %d: incrementConcurrently() = %d, want %d", attempt, got, want)
		}
	}
}
