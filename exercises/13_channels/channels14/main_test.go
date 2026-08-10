package main

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestMergeForwardsValuesAndCloses(t *testing.T) {
	left := make(chan int, 3)
	right := make(chan int, 2)
	for _, value := range []int{1, 3, 5} {
		left <- value
	}
	for _, value := range []int{2, 4} {
		right <- value
	}
	close(left)
	close(right)

	out := merge(left, right)
	var got []int
	deadline := time.After(500 * time.Millisecond)
	for {
		select {
		case value, ok := <-out:
			if !ok {
				sort.Ints(got)
				if want := []int{1, 2, 3, 4, 5}; !reflect.DeepEqual(got, want) {
					t.Fatalf("merge() = %v, want %v", got, want)
				}
				return
			}
			got = append(got, value)
		case <-deadline:
			t.Fatal("merge() did not close its output")
		}
	}
}

func TestMergeWithNoInputsClosesOutput(t *testing.T) {
	out := merge()
	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("merge() returned an unexpected value")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("merge() did not close its empty output")
	}
}
