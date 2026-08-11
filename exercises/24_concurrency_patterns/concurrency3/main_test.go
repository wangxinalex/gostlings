package main

import "testing"

func TestIncrementConcurrentlyReturnsExactCount(t *testing.T) {
	const workers, increments = 20, 1000
	want := int64(workers * increments)
	for attempt := 0; attempt < 3; attempt++ {
		if got := incrementConcurrently(workers, increments); got != want {
			t.Fatalf("attempt %d: got %d, want %d", attempt, got, want)
		}
	}
}
