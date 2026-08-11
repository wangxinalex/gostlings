package main

import (
	"context"
	"testing"
	"time"
)

func TestBufferedPipelineUsesBufferAndCloses(t *testing.T) {
	in := make(chan int, 2)
	in <- 2
	in <- 3
	close(in)
	out := bufferedPipeline(context.Background(), in, 2)
	for i, want := range []int{4, 6} {
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
			t.Fatal("extra output")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("output did not close")
	}
}

func TestBufferedPipelineStopsOnCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	out := bufferedPipeline(ctx, make(chan int), 1)
	cancel()
	select {
	case _, ok := <-out:
		if ok {
			for range out {
			}
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("pipeline did not close")
	}
}
