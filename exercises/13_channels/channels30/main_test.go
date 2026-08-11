package main

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestRunBoundedProcessesEveryJob(t *testing.T) {
	got := runBounded(2, 1, []int{1, 2, 3, 4})
	sort.Ints(got)
	if want := []int{1, 4, 9, 16}; !reflect.DeepEqual(got, want) {
		t.Fatalf("runBounded() = %v, want %v", got, want)
	}
}

func TestRunBoundedWithNoJobsReturns(t *testing.T) {
	if got := runBounded(2, 1, nil); len(got) != 0 {
		t.Fatalf("runBounded() = %v, want no results", got)
	}
}

func TestRunBoundedUsesTheRequestedWorkersAndQueueCapacity(t *testing.T) {
	previousProcess := processBoundedJob
	previousQueue := onBoundedQueue
	started := make(chan struct{}, 2)
	queueCapacity := make(chan int, 1)
	release := make(chan struct{})
	processBoundedJob = func(value int) int {
		started <- struct{}{}
		<-release
		return value * value
	}
	onBoundedQueue = func(capacity int) { queueCapacity <- capacity }
	t.Cleanup(func() {
		processBoundedJob = previousProcess
		onBoundedQueue = previousQueue
	})

	returned := make(chan []int, 1)
	go func() { returned <- runBounded(2, 1, []int{1, 2, 3, 4}) }()
	select {
	case got := <-queueCapacity:
		if got != 1 {
			t.Fatalf("job queue capacity = %d, want 1", got)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runBounded() did not create its bounded job queue")
	}
	for worker := 0; worker < 2; worker++ {
		select {
		case <-started:
		case <-time.After(500 * time.Millisecond):
			t.Fatal("runBounded() did not start every requested worker")
		}
	}
	close(release)
	select {
	case got := <-returned:
		sort.Ints(got)
		if want := []int{1, 4, 9, 16}; !reflect.DeepEqual(got, want) {
			t.Fatalf("runBounded() = %v, want %v", got, want)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runBounded() did not make progress after workers were released")
	}
}
