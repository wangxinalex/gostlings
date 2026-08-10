package main

import (
	"reflect"
	"testing"
	"time"
)

func TestPipelineComposesStagesAndCloses(t *testing.T) {
	in := make(chan int, 3)
	in <- 1
	in <- 2
	in <- 3
	close(in)

	out := pipeline(in)
	var got []int
	deadline := time.After(500 * time.Millisecond)
	for {
		select {
		case value, ok := <-out:
			if !ok {
				if want := []int{3, 5, 7}; !reflect.DeepEqual(got, want) {
					t.Fatalf("pipeline() = %v, want %v", got, want)
				}
				return
			}
			got = append(got, value)
		case <-deadline:
			t.Fatal("pipeline() did not close its output")
		}
	}
}
