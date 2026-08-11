package main

import (
	"reflect"
	"testing"
	"time"
)

func TestRunOrderedRestoresInputOrder(t *testing.T) {
	got := runOrdered(3, []int{4, 1, 3, 2})
	if want := []int{16, 1, 9, 4}; !reflect.DeepEqual(got, want) {
		t.Fatalf("runOrdered() = %v, want %v", got, want)
	}
}

func TestRunOrderedWithNoJobsReturns(t *testing.T) {
	if got := runOrdered(3, nil); len(got) != 0 {
		t.Fatalf("runOrdered() = %v, want no results", got)
	}
}

func TestRunOrderedUsesEveryRequestedWorker(t *testing.T) {
	previous := processOrderedJob
	started := make(chan int, 3)
	completed := make(chan int, 3)
	release := map[int]chan struct{}{
		4: make(chan struct{}),
		1: make(chan struct{}),
		3: make(chan struct{}),
	}
	processOrderedJob = func(value int) int {
		started <- value
		<-release[value]
		completed <- value
		return value * value
	}
	t.Cleanup(func() { processOrderedJob = previous })

	returned := make(chan []int, 1)
	go func() { returned <- runOrdered(3, []int{4, 1, 3}) }()
	for worker := 0; worker < 3; worker++ {
		select {
		case <-started:
		case <-time.After(500 * time.Millisecond):
			t.Fatal("runOrdered() did not start every requested worker")
		}
	}
	for _, value := range []int{3, 1, 4} {
		close(release[value])
		select {
		case got := <-completed:
			if got != value {
				t.Fatalf("processOrderedJob() completed %d, want %d", got, value)
			}
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("processOrderedJob(%d) did not complete after release", value)
		}
	}
	select {
	case got := <-returned:
		if want := []int{16, 1, 9}; !reflect.DeepEqual(got, want) {
			t.Fatalf("runOrdered() = %v, want %v", got, want)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runOrdered() did not finish after workers were released")
	}
}
