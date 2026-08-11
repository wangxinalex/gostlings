package main

import (
	"sync"
	"testing"
)

func TestRunEachCountsActualVisits(t *testing.T) {
	var mu sync.Mutex
	visited := make(map[int]int)
	got := runEach([]int{3, 5, 8}, func(job int) {
		mu.Lock()
		visited[job]++
		mu.Unlock()
	})

	if got != 3 {
		t.Fatalf("runEach() = %d completed visits, want 3", got)
	}
	for _, job := range []int{3, 5, 8} {
		if visited[job] != 1 {
			t.Fatalf("job %d was visited %d times, want 1", job, visited[job])
		}
	}
}
