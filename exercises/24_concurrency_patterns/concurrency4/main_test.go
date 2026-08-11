package main

import (
	"context"
	"testing"
	"time"
)

func TestPipelineRunsBothStagesAndCloses(t *testing.T) {
	in := make(chan int, 3)
	for _, value := range []int{1, 2, 3} {
		in <- value
	}
	close(in)
	out := pipeline(context.Background(), in)
	for i, want := range []int{3, 5, 7} {
		select {
		case got := <-out:
			if got != want {
				t.Fatalf("output %d = %d, want %d", i, got, want)
			}
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("output %d did not arrive", i)
		}
	}
	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("pipeline output has extra values")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("pipeline output did not close")
	}
}

func TestPipelineClosesWhenCanceledWhileOutputIsBlocked(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	in := make(chan int, 1)
	in <- 10
	out := pipeline(ctx, in)
	cancel()
	select {
	case _, ok := <-out:
		if ok {
			for range out {
			}
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("pipeline did not close after cancellation")
	}
}
