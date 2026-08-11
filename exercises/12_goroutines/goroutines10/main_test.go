package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestReviewRunReviewsSeveralJobs(t *testing.T) {
	got := reviewRun([]int{6, 2, 9})
	want := []string{"reviewed 2", "reviewed 6", "reviewed 9"}
	sort.Strings(got)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("reviewRun() = %v, want %v", got, want)
	}
}

func TestReviewRunHandlesNoJobs(t *testing.T) {
	if got := reviewRun(nil); len(got) != 0 {
		t.Fatalf("reviewRun(nil) returned %d results, want 0", len(got))
	}
}
