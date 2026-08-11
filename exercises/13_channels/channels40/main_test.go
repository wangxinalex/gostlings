package main

import (
	"reflect"
	"testing"
	"time"
)

func TestRelayForwardsValuesThenCloses(t *testing.T) {
	in := make(chan int, 2)
	in <- 4
	in <- 9
	close(in)
	if got := collect40(t, relay(make(chan struct{}), in)); !reflect.DeepEqual(got, []int{4, 9}) {
		t.Fatalf("relay() = %v, want [4 9]", got)
	}
}
func TestRelayStopsWhileWaitingForInput(t *testing.T) {
	stop, in := make(chan struct{}), make(chan int)
	out := relay(stop, in)
	close(stop)
	closed40(t, out)
}
func TestRelayStopsWhileDownstreamIsBlocked(t *testing.T) {
	stop, in := make(chan struct{}), make(chan int)
	out := relay(stop, in)
	sent := make(chan struct{})
	go func() { in <- 7; close(sent) }()
	select {
	case <-sent:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("relay() did not receive its input")
	}
	close(stop)
	closed40(t, out)
}
func collect40(t *testing.T, out <-chan int) []int {
	t.Helper()
	var got []int
	for {
		select {
		case v, ok := <-out:
			if !ok {
				return got
			}
			got = append(got, v)
		case <-time.After(500 * time.Millisecond):
			t.Fatal("relay() did not close")
		}
	}
}
func closed40(t *testing.T, out <-chan int) {
	t.Helper()
	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("relay() sent after cancellation")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("relay() did not close after cancellation")
	}
}
