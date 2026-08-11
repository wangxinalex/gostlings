package main

import (
	"context"
	"testing"
	"time"
)

func TestBatchFlushesBySizeAndExplicitEvent(t *testing.T) {
	in := make(chan int, 3)
	flush := make(chan time.Time, 1)
	out := batch(context.Background(), in, flush, 2)
	in <- 1
	in <- 2
	select {
	case got := <-out:
		if len(got) != 2 || got[0] != 1 || got[1] != 2 {
			t.Fatalf("full batch = %v", got)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("full batch did not arrive")
	}
	in <- 3
	flush <- time.Now()
	select {
	case got := <-out:
		if len(got) != 1 || got[0] != 3 {
			t.Fatalf("flushed batch = %v", got)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("flushed batch did not arrive")
	}
	close(in)
	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("batch output had unexpected values")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("batch output did not close")
	}
}

func TestBatchDiscardsPartialBatchOnCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	in := make(chan int, 1)
	flush := make(chan time.Time)
	out := batch(ctx, in, flush, 3)
	in <- 9
	cancel()
	select {
	case value, ok := <-out:
		if ok {
			t.Fatalf("canceled batch emitted %v", value)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("batch did not close after cancellation")
	}
}
