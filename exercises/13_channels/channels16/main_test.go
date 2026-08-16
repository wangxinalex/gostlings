package main

import (
	"testing"
	"time"
)

func TestProduceSendsValuesBeforeCancellation(t *testing.T) {
	stop := make(chan struct{})
	out := produce(stop)

	for want := 1; want <= 2; want++ {
		select {
		case got := <-out:
			if got != want {
				t.Fatalf("produce() value = %d, want %d", got, want)
			}
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("produce() did not send value %d", want)
		}
	}

	close(stop)
	waitForProduceClose(t, out)
}

func TestProduceStopsWhenNoReceiverIsAvailable(t *testing.T) {
	stop := make(chan struct{})
	out := produce(stop)
	close(stop)

	waitForProduceClose(t, out)
}

func waitForProduceClose(t *testing.T, out <-chan int) {
	t.Helper()
	timer := time.NewTimer(500 * time.Millisecond)
	defer timer.Stop()

	// Cancellation makes the stop case ready, but a concurrent send may also
	// be ready. Select does not prioritize stop, so discard any in-flight
	// values and assert that the producer eventually closes out.
	for {
		select {
		case _, ok := <-out:
			if !ok {
				return
			}
		case <-timer.C:
			t.Fatal("produce() did not close after cancellation")
		}
	}
}
