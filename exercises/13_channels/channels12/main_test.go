package main

import (
	"reflect"
	"testing"
	"time"
)

func TestRelayForwardsInputThenClosesOutput(t *testing.T) {
	in := make(chan int, 2)
	in <- 4
	in <- 9
	close(in)

	out := relay(make(chan struct{}), in)
	var got []int
	for {
		select {
		case value, ok := <-out:
			if !ok {
				want := []int{4, 9}
				if !reflect.DeepEqual(got, want) {
					t.Fatalf("relay() values = %v, want %v", got, want)
				}
				return
			}
			got = append(got, value)
		case <-time.After(500 * time.Millisecond):
			t.Fatal("relay() did not close after its input closed")
		}
	}
}

func TestRelayStopsWhileOutputSendIsBlocked(t *testing.T) {
	stop := make(chan struct{})
	in := make(chan int)
	out := relay(stop, in)
	sent := make(chan struct{})
	go func() {
		in <- 7
		close(sent)
	}()

	select {
	case <-sent:
		// The relay received the input, so it can now only be blocked sending it.
	case <-time.After(500 * time.Millisecond):
		t.Fatal("relay() did not receive its input")
	}

	close(stop)
	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("relay() sent a value after cancellation")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("relay() stayed blocked sending after cancellation")
	}
}
