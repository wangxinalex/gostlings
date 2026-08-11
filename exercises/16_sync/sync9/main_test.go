package main

import (
	"testing"
	"time"
)

func TestPopWaitsUntilPushSignals(t *testing.T) {
	q := newIntQueue()
	result := make(chan int, 1)
	go func() { result <- q.Pop() }()
	select {
	case value := <-result:
		t.Fatalf("Pop() returned %d before a value was pushed", value)
	case <-time.After(20 * time.Millisecond):
	}
	q.Push(42)
	select {
	case value := <-result:
		if value != 42 {
			t.Fatalf("Pop() = %d, want 42", value)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("Pop() did not resume after Push")
	}
}

func TestQueuePreservesFIFOOrder(t *testing.T) {
	q := newIntQueue()
	q.Push(1)
	q.Push(2)
	if got := q.Pop(); got != 1 {
		t.Fatalf("first Pop() = %d, want 1", got)
	}
	if got := q.Pop(); got != 2 {
		t.Fatalf("second Pop() = %d, want 2", got)
	}
}
