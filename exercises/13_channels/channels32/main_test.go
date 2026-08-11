package main

import (
	"testing"
	"time"
)

func TestServeRepliesOnEachRequestChannelAndStopsAfterInputCloses(t *testing.T) {
	requests := make(chan request)
	done := serve(requests)
	first, second := make(chan int), make(chan int)
	go func() {
		requests <- request{value: 3, reply: first}
		requests <- request{value: 7, reply: second}
		close(requests)
	}()
	if got := receive32(t, first); got != 6 {
		t.Fatalf("first reply = %d, want 6", got)
	}
	if got := receive32(t, second); got != 14 {
		t.Fatalf("second reply = %d, want 14", got)
	}
	select {
	case <-done:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("serve() did not finish after requests closed")
	}
}
func receive32(t *testing.T, reply <-chan int) int {
	t.Helper()
	select {
	case value := <-reply:
		return value
	case <-time.After(500 * time.Millisecond):
		t.Fatal("serve() did not send a reply")
		return 0
	}
}
