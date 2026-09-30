package main

import (
	"testing"
	"time"
)

func TestSumParallelAddsAnEvenSplit(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8}
	if got := sumParallel(values, 4); got != 36 {
		t.Fatalf("sumParallel(%v, 4) = %d, want 36", values, got)
	}
}

func TestSumParallelAddsAnUnevenSplit(t *testing.T) {
	values := []int{3, 8, 1, 9, 4, 7, 2}
	if got := sumParallel(values, 3); got != 34 {
		t.Fatalf("sumParallel(%v, 3) = %d, want 34", values, got)
	}
}

func TestSumParallelHandlesFewerValuesThanWorkers(t *testing.T) {
	completed := make(chan int, 1)
	go func() { completed <- sumParallel([]int{4, 5}, 8) }()

	select {
	case got := <-completed:
		if got != 9 {
			t.Fatalf("sumParallel([4 5], 8) = %d, want 9", got)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("sumParallel did not join more workers than values")
	}
}

func TestSumParallelHandlesEmptyInput(t *testing.T) {
	if got := sumParallel(nil, 4); got != 0 {
		t.Fatalf("sumParallel(nil, 4) = %d, want 0", got)
	}
	if got := sumParallel([]int{}, 3); got != 0 {
		t.Fatalf("sumParallel([], 3) = %d, want 0", got)
	}
}
