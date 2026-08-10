package main

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestParallelHonorsConcurrencyLimit(t *testing.T) {
	var mu sync.Mutex
	active, maxActive := 0, 0
	work := func(value int) int {
		mu.Lock()
		active++
		if active > maxActive {
			maxActive = active
		}
		mu.Unlock()

		time.Sleep(5 * time.Millisecond)

		mu.Lock()
		active--
		mu.Unlock()
		return value * value
	}

	got := parallel(2, []int{1, 2, 3, 4, 5, 6}, work)
	if want := []int{1, 4, 9, 16, 25, 36}; !reflect.DeepEqual(got, want) {
		t.Fatalf("parallel() = %v, want %v", got, want)
	}
	if maxActive > 2 {
		t.Fatalf("parallel() ran %d jobs at once, limit is 2", maxActive)
	}
}

func TestParallelWithNoJobsReturns(t *testing.T) {
	if got := parallel(2, nil, func(value int) int { return value }); len(got) != 0 {
		t.Fatalf("parallel() with no jobs = %v, want empty result", got)
	}
}
