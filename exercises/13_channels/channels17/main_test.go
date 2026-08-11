package main

import (
	"reflect"
	"testing"
	"time"
)

func TestForwardForwardsInputThenClosesOutput(t *testing.T) {
	in := make(chan int, 2)
	in <- 3
	in <- 8
	close(in)

	out := forward(make(chan struct{}), in)
	var got []int
	for {
		select {
		case value, ok := <-out:
			if !ok {
				want := []int{3, 8}
				if !reflect.DeepEqual(got, want) {
					t.Fatalf("forward() values = %v, want %v", got, want)
				}
				return
			}
			got = append(got, value)
		case <-time.After(500 * time.Millisecond):
			t.Fatal("forward() did not close after its input closed")
		}
	}
}

func TestForwardStopsWhileWaitingForInput(t *testing.T) {
	stop := make(chan struct{})
	out := forward(stop, make(chan int))
	close(stop)
	waitForForwardClose(t, out)
}

func TestForwardStopsWhileOutputSendIsBlocked(t *testing.T) {
	stop := make(chan struct{})
	in := make(chan int)
	out := forward(stop, in)
	sent := make(chan struct{})
	go func() {
		in <- 6
		close(sent)
	}()

	select {
	case <-sent:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("forward() did not receive its input")
	}

	close(stop)
	waitForForwardClose(t, out)
}

func waitForForwardClose(t *testing.T, out <-chan int) {
	t.Helper()
	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("forward() sent a value after cancellation")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("forward() did not close after cancellation")
	}
}
