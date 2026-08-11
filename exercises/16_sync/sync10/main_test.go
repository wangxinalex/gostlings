package main

import (
	"testing"
	"time"
)

func TestCloseWakesEveryBlockedPop(t *testing.T) {
	q := newClosableQueue()
	results := make(chan bool, 2)
	for i := 0; i < 2; i++ {
		go func() {
			_, ok := q.Pop()
			results <- ok
		}()
	}
	select {
	case <-results:
		t.Fatal("Pop() returned before queue close")
	case <-time.After(20 * time.Millisecond):
	}
	q.Close()
	for i := 0; i < 2; i++ {
		select {
		case ok := <-results:
			if ok {
				t.Fatal("closed empty queue returned a value")
			}
		case <-time.After(500 * time.Millisecond):
			t.Fatal("Close() did not wake every Pop")
		}
	}
}

func TestClosableQueueDrainsValuesAndRejectsLaterPushes(t *testing.T) {
	q := newClosableQueue()
	if !q.Push(9) {
		t.Fatal("Push() rejected an open queue")
	}
	q.Close()
	if value, ok := q.Pop(); !ok || value != 9 {
		t.Fatalf("Pop() = (%d, %v), want queued value", value, ok)
	}
	if q.Push(10) {
		t.Fatal("Push() accepted a closed queue")
	}
	if _, ok := q.Pop(); ok {
		t.Fatal("Pop() returned ok after the queue drained")
	}
}
