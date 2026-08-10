package main

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestSquareWorkersForwardsAllJobsAndCloses(t *testing.T) {
	jobs := make(chan int, 4)
	for _, job := range []int{1, 2, 3, 4} {
		jobs <- job
	}
	close(jobs)

	out := squareWorkers(2, jobs)
	var got []int
	deadline := time.After(500 * time.Millisecond)
	for {
		select {
		case value, ok := <-out:
			if !ok {
				sort.Ints(got)
				want := []int{1, 4, 9, 16}
				if !reflect.DeepEqual(got, want) {
					t.Fatalf("squareWorkers() = %v, want %v", got, want)
				}
				return
			}
			got = append(got, value)
		case <-deadline:
			t.Fatal("squareWorkers() did not close its output")
		}
	}
}
