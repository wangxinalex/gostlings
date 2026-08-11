package main

import (
	"testing"
	"time"
)

func TestTrySendDeliversToAvailableBuffer(t *testing.T) {
	ch := make(chan int, 1)
	if ok := trySend(ch, 7); !ok {
		t.Fatal("trySend() = false, want true when buffer has room")
	}

	select {
	case got := <-ch:
		if got != 7 {
			t.Fatalf("channel received %d, want 7", got)
		}
	default:
		t.Fatal("trySend() returned true without sending a value")
	}
}

func TestTrySendRejectsFullOrUnreadyChannel(t *testing.T) {
	t.Run("full buffer", func(t *testing.T) {
		ch := make(chan int, 1)
		ch <- 3
		if ok := trySendWithin(t, ch, 7); ok {
			t.Fatal("trySend() = true, want false when buffer is full")
		}
		if got := <-ch; got != 3 {
			t.Fatalf("full channel contained %d after trySend, want 3", got)
		}
	})

	t.Run("unbuffered channel without receiver", func(t *testing.T) {
		if ok := trySendWithin(t, make(chan int), 7); ok {
			t.Fatal("trySend() = true, want false without a ready receiver")
		}
	})
}

func trySendWithin(t *testing.T, ch chan<- int, value int) bool {
	t.Helper()
	result := make(chan bool, 1)
	go func() {
		result <- trySend(ch, value)
	}()

	select {
	case ok := <-result:
		return ok
	case <-time.After(500 * time.Millisecond):
		t.Fatal("trySend() blocked instead of returning immediately")
		return false
	}
}
