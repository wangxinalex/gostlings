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
	previousStart := onBoundedProcessStart
	started := make(chan int, 3)
	processing := make(chan int, 3)
	queueCapacity := make(chan int, 1)
	release := make(chan struct{})
	processBoundedJob = func(value int) int {
		started <- value
		<-release
		return value * value
	}
	onBoundedQueue = func(capacity int) { queueCapacity <- capacity }
	onBoundedProcessStart = func(value int) { processing <- value }
	t.Cleanup(func() {
		processBoundedJob = previousProcess
		onBoundedQueue = previousQueue
		onBoundedProcessStart = previousStart
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
	select {
	case value := <-processing:
		_ = value
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runBounded() did not enter processing")
	}
	select {
	case value := <-processing:
		_ = value
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runBounded() did not enter processing with both workers")
	}
	select {
	case value := <-processing:
		t.Fatalf("runBounded() started a third concurrent job (%d) before workers were released", value)
	case <-time.After(100 * time.Millisecond):
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
