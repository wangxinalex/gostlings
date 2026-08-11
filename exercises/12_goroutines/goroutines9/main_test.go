package main

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestRunWithEarlyReturnStillJoinsEveryWorker(t *testing.T) {
	completed := make(chan []string, 1)
	go func() {
		completed <- runWithEarlyReturn([]int{1, 2, 3}, 2)
	}()

	select {
	case got := <-completed:
		want := []string{"", "job 1 done", "job 3 done"}
		sort.Strings(got)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("runWithEarlyReturn() = %v, want %v", got, want)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runWithEarlyReturn did not join an early-return worker")
	}
}
