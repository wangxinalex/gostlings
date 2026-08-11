package main

import (
	"reflect"
	"testing"
	"time"
)

func TestProduceSendsItsSequenceThenCloses(t *testing.T) {
	if got := collect42(t, produce(make(chan struct{}))); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Fatalf("produce() = %v, want [1 2 3]", got)
	}
}

func TestProduceStopsWhenConsumerAbandonsOutput(t *testing.T) {
	stop := make(chan struct{})
	out := produce(stop)
	select {
	case got := <-out:
		if got != 1 {
			t.Fatalf("first produce() value = %d, want 1", got)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("produce() did not start")
	}

	close(stop)
	closed42(t, out)
}

func collect42(t *testing.T, out <-chan int) []int {
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
			t.Fatal("produce() did not close")
		}
	}
}

func closed42(t *testing.T, out <-chan int) {
	t.Helper()
	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("produce() sent after stop")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("produce() did not close after stop")
	}
}
