package main

import (
	"context"
	"testing"
	"time"
)

func TestSquareTransformsAndCloses(t *testing.T) {
	in := make(chan int, 3)
	in <- 1
	in <- 2
	in <- 3
	close(in)
	out := square(context.Background(), in)
	got := make([]int, 0, 3)
	for {
		select {
		case value, ok := <-out:
			if !ok {
				if len(got) != 3 || got[0] != 1 || got[1] != 4 || got[2] != 9 {
					t.Fatalf("squares = %v, want [1 4 9]", got)
				}
				return
			}
			got = append(got, value)
		case <-time.After(500 * time.Millisecond):
			t.Fatal("square() did not close output")
		}
	}
}

func TestSquareStopsWhenOutputIsBlocked(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	in := make(chan int, 1)
	in <- 2
	out := square(ctx, in)
	cancel()
	select {
	case _, ok := <-out:
		if ok {
			for range out {
			}
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("square() did not close after cancellation")
	}
}
