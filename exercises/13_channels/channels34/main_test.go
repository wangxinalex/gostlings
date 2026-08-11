package main

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestRunProcessesEveryJob(t *testing.T) {
	got := run(2, []int{1, 2, 3, 4})
	sort.Ints(got)
	if want := []int{1, 4, 9, 16}; !reflect.DeepEqual(got, want) {
		t.Fatalf("run() = %v, want %v", got, want)
	}
}
func TestRunReturnsForEmptyJobs(t *testing.T) {
	if got := run(2, nil); len(got) != 0 {
		t.Fatalf("run(nil) = %v, want no values", got)
	}
}
func TestRunWithZeroWorkersReturnsWithoutDeadlock(t *testing.T) {
	returned := make(chan []int, 1)
	go func() { returned <- run(0, []int{1, 2}) }()
	select {
	case got := <-returned:
		if len(got) != 0 {
			t.Fatalf("run(0, jobs) = %v, want no values", got)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("run(0, jobs) deadlocked")
	}
}
