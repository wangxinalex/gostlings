package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestRunWorkersWaitsForEveryWorker(t *testing.T) {
	got := runWorkers(4)
	if len(got) != 4 {
		t.Fatalf("runWorkers(4) returned %d completions, want 4", len(got))
	}
	want := []string{"worker 0 done", "worker 1 done", "worker 2 done", "worker 3 done"}
	sort.Strings(got)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("runWorkers(4) = %v, want distinct worker completions %v", got, want)
	}
}

func TestRunWorkersHandlesZeroWorkers(t *testing.T) {
	if got := runWorkers(0); len(got) != 0 {
		t.Fatalf("runWorkers(0) returned %d completions, want 0", len(got))
	}
}
