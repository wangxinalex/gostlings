package main

import (
	"reflect"
	"testing"
	"time"
)

func TestCancellablePipelineForwardsNormalInput(t *testing.T) {
	stop := make(chan struct{})
	in := make(chan int, 2)
	in <- 1
	in <- 2
	close(in)

	out := pipeline(stop, in)
	var got []int
	for value := range out {
		got = append(got, value)
	}
	if want := []int{3, 5}; !reflect.DeepEqual(got, want) {
		t.Fatalf("pipeline() = %v, want %v", got, want)
	}
}

func TestCancellablePipelineStopsBlockedInput(t *testing.T) {
	stop := make(chan struct{})
	out := pipeline(stop, make(chan int))
	close(stop)

	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("pipeline() emitted a value after cancellation")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("pipeline() did not close after cancellation")
	}
}
