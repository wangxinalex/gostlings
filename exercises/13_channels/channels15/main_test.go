package main

import (
	"reflect"
	"testing"
	"time"
)

func TestCancellableMergeStopsBlockedInput(t *testing.T) {
	stop := make(chan struct{})
	blocked := make(chan int)
	out := merge(stop, blocked)
	close(stop)

	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("merge() emitted a value after cancellation")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("merge() left a forwarder blocked on input")
	}
}

func TestCancellableMergeForwardsNormalInputs(t *testing.T) {
	stop := make(chan struct{})
	input := make(chan int, 2)
	input <- 4
	input <- 5
	close(input)

	out := merge(stop, input)
	var got []int
	for value := range out {
		got = append(got, value)
	}
	if want := []int{4, 5}; !reflect.DeepEqual(got, want) {
		t.Fatalf("merge() = %v, want %v", got, want)
	}
}
