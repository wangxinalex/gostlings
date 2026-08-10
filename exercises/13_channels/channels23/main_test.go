package main

import (
	"testing"
	"time"
)

func TestRateLimitWaitsForTickBeforeForwarding(t *testing.T) {
	ticks := make(chan time.Time)
	in := make(chan int, 1)
	in <- 7
	close(in)
	out := rateLimit(ticks, in)

	select {
	case value := <-out:
		t.Fatalf("rateLimit() forwarded %d before a tick", value)
	case <-time.After(20 * time.Millisecond):
	}

	ticks <- time.Time{}
	select {
	case value, ok := <-out:
		if !ok || value != 7 {
			t.Fatalf("rateLimit() returned (%d, %v), want (7, true)", value, ok)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("rateLimit() did not forward after a tick")
	}

	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("rateLimit() returned an extra value")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("rateLimit() did not close after input drained")
	}
}
