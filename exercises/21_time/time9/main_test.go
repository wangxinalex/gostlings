package main

import (
	"testing"
	"time"
)

func scriptedClock(times ...time.Time) func() time.Time {
	index := 0
	return func() time.Time {
		if index >= len(times) {
			return times[len(times)-1]
		}
		current := times[index]
		index++
		return current
	}
}

func TestMeasureAcceptsWorkWithinTheBudget(t *testing.T) {
	start := time.Date(2024, time.March, 1, 9, 0, 0, 0, time.UTC)
	elapsed, ok := measure(scriptedClock(start, start.Add(40*time.Millisecond)), func() {})
	if !ok {
		t.Fatalf("measure() ok = false for %v, want true", elapsed)
	}
	if elapsed != 40*time.Millisecond {
		t.Fatalf("measure() elapsed = %v, want %v", elapsed, 40*time.Millisecond)
	}
}

func TestMeasureRejectsWorkPastTheBudget(t *testing.T) {
	start := time.Date(2024, time.March, 1, 9, 0, 0, 0, time.UTC)
	elapsed, ok := measure(scriptedClock(start, start.Add(measureBudget+time.Millisecond)), func() {})
	if ok {
		t.Fatalf("measure() ok = true for %v, want false", elapsed)
	}
	if elapsed != measureBudget+time.Millisecond {
		t.Fatalf("measure() elapsed = %v, want %v", elapsed, measureBudget+time.Millisecond)
	}
}

func TestMeasureTreatsTheBudgetItselfAsWithinBudget(t *testing.T) {
	start := time.Date(2024, time.March, 1, 9, 0, 0, 0, time.UTC)
	elapsed, ok := measure(scriptedClock(start, start.Add(measureBudget)), func() {})
	if elapsed != measureBudget || !ok {
		t.Fatalf("measure() = (%v, %v), want (%v, true)", elapsed, ok, measureBudget)
	}
}

func TestMeasureReportsTheExactElapsedTime(t *testing.T) {
	current := time.Date(2024, time.March, 1, 9, 0, 0, 0, time.UTC)
	elapsed, ok := measure(func() time.Time { return current }, func() {
		current = current.Add(250 * time.Millisecond)
	})
	if elapsed != 250*time.Millisecond {
		t.Fatalf("measure() elapsed = %v, want %v", elapsed, 250*time.Millisecond)
	}
	if ok {
		t.Fatalf("measure() ok = true for %v, want false", elapsed)
	}
}
