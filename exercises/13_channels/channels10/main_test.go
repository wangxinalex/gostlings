package main

import (
	"reflect"
	"testing"
	"time"
)

func TestProduceStopsWhenCancellationArrivesBeforeReceive(t *testing.T) {
	stop := make(chan struct{})
	out := produce(stop)
	close(stop)

	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("produce() emitted a value after cancellation")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("produce() did not stop after cancellation")
	}
}

func TestProduceClosesAfterNormalValues(t *testing.T) {
	out := produce(make(chan struct{}))
	var got []int
	for value := range out {
		got = append(got, value)
	}

	if want := []int{1, 2, 3}; !reflect.DeepEqual(got, want) {
		t.Fatalf("produce() = %v, want %v", got, want)
	}
}
