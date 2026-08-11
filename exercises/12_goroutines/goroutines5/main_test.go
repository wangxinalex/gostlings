package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestRunJobsWaitsForDynamicInput(t *testing.T) {
	got := runJobs([]int{4, 9, 2}, func(job int) string {
		return fmt.Sprintf("job %d", job)
	})
	want := []string{"job 2", "job 4", "job 9"}
	sort.Strings(got)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("runJobs() = %v, want %v", got, want)
	}
}

func TestRunJobsHandlesNoJobs(t *testing.T) {
	if got := runJobs(nil, func(int) string { return "unexpected" }); len(got) != 0 {
		t.Fatalf("runJobs(nil) returned %d results, want 0", len(got))
	}
}
