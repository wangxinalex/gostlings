package main

import (
	"context"
	"testing"
	"time"
)

func TestServeProcessesAndClosesInOrder(t *testing.T) {
	jobs := make(chan int, 2)
	jobs <- 2
	jobs <- 3
	close(jobs)
	results, done := serve(context.Background(), jobs)
	for i, want := range []int{4, 6} {
		select {
		case got := <-results:
			if got != want {
				t.Fatalf("result %d = %d, want %d", i, got, want)
			}
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("result %d did not arrive", i)
		}
	}
	select {
	case <-done:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("serve() did not signal done")
	}
}

func TestServeStopsAcceptingAfterCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	jobs := make(chan int)
	results, done := serve(ctx, jobs)
	cancel()
	select {
	case <-done:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("serve() did not stop")
	}
	select {
	case _, ok := <-results:
		if ok {
			t.Fatal("results remained open after shutdown")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("results did not close after shutdown")
	}
}
