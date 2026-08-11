package main

import (
	"testing"
	"time"
)

func TestProduceSendsAValueBeforeCancellation(t *testing.T) {
	stop := make(chan struct{})
	out := produce(stop)

	select {
	case got := <-out:
		if got != 1 {
			t.Fatalf("produce() first value = %d, want 1", got)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("produce() did not send its first value")
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
	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("produce() sent a value after cancellation")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("produce() did not close after cancellation")
	}
}
