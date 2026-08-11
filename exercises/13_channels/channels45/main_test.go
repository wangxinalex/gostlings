package main

import (
	"testing"
	"time"
)

func TestRateLimitRequiresOneTokenForEachInputAndCloses(t *testing.T) {
	tokens := make(chan struct{})
	in := make(chan int, 2)
	in <- 4
	in <- 9
	close(in)
	out := rateLimit(tokens, in)
	select {
	case got := <-out:
		t.Fatalf("rateLimit() forwarded %d without a token", got)
	default:
	}

	firstToken := make(chan struct{})
	go func() { tokens <- struct{}{}; close(firstToken) }()
	if got := receive45(t, out); got != 4 {
		t.Fatalf("first rateLimit() value = %d, want 4", got)
	}
	select {
	case <-firstToken:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("rateLimit() did not consume its first token")
	}
	secondToken := make(chan struct{})
	go func() { tokens <- struct{}{}; close(secondToken) }()
	if got := receive45(t, out); got != 9 {
		t.Fatalf("second rateLimit() value = %d, want 9", got)
	}
	select {
	case <-secondToken:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("rateLimit() did not consume its second token")
	}
	closed45(t, out)
}

func receive45(t *testing.T, out <-chan int) int {
	t.Helper()
	select {
	case value := <-out:
		return value
	case <-time.After(500 * time.Millisecond):
		t.Fatal("rateLimit() did not forward after a token")
		return 0
	}
}

func closed45(t *testing.T, out <-chan int) {
	t.Helper()
	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("rateLimit() sent after input closed")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("rateLimit() did not close after input closed")
	}
}
