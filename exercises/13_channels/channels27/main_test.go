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
	started := make(chan struct{}, 3)
	release := make(chan struct{})
	processOrderedJob = func(value int) int {
		started <- struct{}{}
		<-release
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
	close(release)
	select {
	case got := <-returned:
		if want := []int{16, 1, 9}; !reflect.DeepEqual(got, want) {
			t.Fatalf("runOrdered() = %v, want %v", got, want)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runOrdered() did not finish after workers were released")
	}
}
