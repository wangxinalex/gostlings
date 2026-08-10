package main

import (
	"reflect"
	"testing"
	"time"
)

func TestSquareStageTransformsAndCloses(t *testing.T) {
	in := make(chan int, 3)
	in <- 1
	in <- 2
	in <- 3
	close(in)

	out := square(in)
	var got []int
	deadline := time.After(500 * time.Millisecond)
	for {
		select {
		case value, ok := <-out:
			if !ok {
				if want := []int{1, 4, 9}; !reflect.DeepEqual(got, want) {
					t.Fatalf("square() = %v, want %v", got, want)
				}
				return
			}
			got = append(got, value)
		case <-deadline:
			t.Fatal("square() did not close its output")
		}
	}
}
