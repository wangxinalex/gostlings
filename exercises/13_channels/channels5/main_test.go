package main

import (
	"reflect"
	"testing"
)

func TestDrainClosedReturnsValuesBeforeClosedState(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 4
	ch <- 0
	ch <- 9
	close(ch)

	got := drainClosed(ch)
	want := []int{4, 0, 9}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("drainClosed() = %v, want %v", got, want)
	}
}
