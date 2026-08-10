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
	left <- 1
	left <- 3
	left <- 5
	right <- 2
	right <- 4
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
				want := []int{1, 2, 3, 4, 5}
				if !reflect.DeepEqual(got, want) {
					t.Fatalf("values = %v, want %v", got, want)
				}
				return
			}
			got = append(got, value)
		case <-deadline:
			t.Fatal("merge did not forward all values and close its output")
		}
	}
}

func TestMergeWithNoInputsClosesOutput(t *testing.T) {
	out := merge()
	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("merge() returned a value for empty input")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("merge() did not close its output for empty input")
	}
}
