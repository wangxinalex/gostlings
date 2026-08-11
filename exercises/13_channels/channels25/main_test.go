package main

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestMergeClosesForNoInputs(t *testing.T) {
	if got := collectRobustMerge(t, merge()); len(got) != 0 {
		t.Fatalf("merge() = %v, want no values", got)
	}
}

func TestMergeDrainsClosedAndBufferedInputs(t *testing.T) {
	alreadyClosed := make(chan int)
	close(alreadyClosed)
	bufferedInput := robustBuffered(1, 2, 3)
	got := collectRobustMerge(t, merge(alreadyClosed, bufferedInput))
	if want := []int{1, 2, 3}; !reflect.DeepEqual(got, want) {
		t.Fatalf("merge() = %v, want %v", got, want)
	}
}

func TestMergeWaitsForEveryNonNilInput(t *testing.T) {
	blocked := make(chan int)
	out := merge(robustBuffered(1), blocked)
	sent := make(chan struct{})
	go func() {
		blocked <- 2
		close(blocked)
		close(sent)
	}()
	got := collectRobustMerge(t, out)
	select {
	case <-sent:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("merge() did not wait for its blocked input")
	}
	sort.Ints(got)
	if want := []int{1, 2}; !reflect.DeepEqual(got, want) {
		t.Fatalf("merge() = %v, want %v", got, want)
	}
}

func robustBuffered(values ...int) <-chan int {
	in := make(chan int, len(values))
	for _, value := range values {
		in <- value
	}
	close(in)
	return in
}

func collectRobustMerge(t *testing.T, out <-chan int) []int {
	t.Helper()
	var got []int
	for {
		select {
		case value, ok := <-out:
			if !ok {
				return got
			}
			got = append(got, value)
		case <-time.After(500 * time.Millisecond):
			t.Fatal("merge() did not close")
		}
	}
}
