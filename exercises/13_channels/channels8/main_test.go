package main

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestRunProcessesEveryJob(t *testing.T) {
	result := make(chan []int, 1)
	go func() {
		result <- run(3, []int{1, 2, 3, 4, 5})
	}()

	select {
	case got := <-result:
		sort.Ints(got)
		want := []int{1, 4, 9, 16, 25}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("run() = %v, want %v", got, want)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("run() did not finish processing all jobs")
	}
}

func TestRunWithNoJobsReturns(t *testing.T) {
	result := make(chan []int, 1)
	go func() {
		result <- run(2, nil)
	}()

	select {
	case got := <-result:
		if len(got) != 0 {
			t.Fatalf("run() with no jobs = %v, want no results", got)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("run() did not finish with no jobs")
	}
}
