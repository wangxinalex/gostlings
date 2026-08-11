package main

import (
	"testing"
	"time"
)

func TestRunTasksRunsEveryDynamicJob(t *testing.T) {
	previous := runTask
	called := make(chan int, 4)
	runTask = func(job int) int { called <- job; return job * 2 }
	t.Cleanup(func() { runTask = previous })
	jobs := []int{3, 1, 4, 2}
	if got := runTasks(jobs); got != 20 {
		t.Fatalf("runTasks() = %d, want 20", got)
	}
	seen := make(map[int]bool)
	for range jobs {
		select {
		case job := <-called:
			seen[job] = true
		case <-time.After(500 * time.Millisecond):
			t.Fatal("runTasks() did not run every job")
		}
	}
	for _, job := range jobs {
		if !seen[job] {
			t.Fatalf("job %d was not run", job)
		}
	}
}

func TestRunTasksHandlesEmptyInput(t *testing.T) {
	if got := runTasks(nil); got != 0 {
		t.Fatalf("runTasks(nil) = %d, want 0", got)
	}
}
