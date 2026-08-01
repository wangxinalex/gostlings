package main

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestSquareProducesValuesAndCloses(t *testing.T) {
	in := make(chan int, 3)
	in <- 1
	in <- 2
	in <- 3
	close(in)

	out := square(context.Background(), in)
	var got []int
	for {
		select {
		case value, ok := <-out:
			if !ok {
				if !reflect.DeepEqual(got, []int{1, 4, 9}) {
					t.Fatalf("values = %v, want %v", got, []int{1, 4, 9})
				}
				return
			}
			got = append(got, value)
		case <-time.After(500 * time.Millisecond):
			t.Fatal("square did not close its output")
		}
	}
}

func TestSquareStopsWhileBlockedOnSend(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	in := make(chan int, 1)
	in <- 2
	out := square(ctx, in)
	cancel()

	deadline := time.After(500 * time.Millisecond)
	for {
		select {
		case _, ok := <-out:
			if !ok {
				return
			}
		case <-deadline:
			t.Fatal("square did not close its output after cancellation")
		}
	}
}
