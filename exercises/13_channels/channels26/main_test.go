package main

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestRunProcessesEveryJob(t *testing.T) {
	got := run(3, []int{1, 2, 3, 4})
	sort.Ints(got)
	if want := []int{1, 4, 9, 16}; !reflect.DeepEqual(got, want) {
		t.Fatalf("run() = %v, want %v", got, want)
	}
}

func TestRunWithNoJobsReturns(t *testing.T) {
	if got := run(3, nil); len(got) != 0 {
		t.Fatalf("run() = %v, want no results", got)
	}
}

func TestRunUsesEveryRequestedWorker(t *testing.T) {
	previous := processJob
	started := make(chan struct{}, 3)
	release := make(chan struct{})
	processJob = func(value int) int {
		started <- struct{}{}
		<-release
		return value * value
	}
	t.Cleanup(func() { processJob = previous })

	returned := make(chan []int, 1)
	go func() { returned <- run(3, []int{1, 2, 3}) }()
	for worker := 0; worker < 3; worker++ {
		select {
		case <-started:
		case <-time.After(500 * time.Millisecond):
			t.Fatal("run() did not start every requested worker")
		}
	}
	close(release)
	select {
	case got := <-returned:
		sort.Ints(got)
		if want := []int{1, 4, 9}; !reflect.DeepEqual(got, want) {
			t.Fatalf("run() = %v, want %v", got, want)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("run() did not finish after workers were released")
	}
}

func TestRunWaitsForEveryWorkerExitBeforeReturning(t *testing.T) {
	previous := onWorkerExit
	exited := make(chan struct{}, 3)
	release := make(chan struct{})
	onWorkerExit = func() {
		exited <- struct{}{}
		<-release
	}
	t.Cleanup(func() { onWorkerExit = previous })

	returned := make(chan []int, 1)
	go func() { returned <- run(3, []int{1, 2, 3}) }()
	for worker := 0; worker < 3; worker++ {
		select {
		case <-exited:
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("worker %d did not reach its exit hook", worker+1)
		}
	}
	select {
	case got := <-returned:
		t.Fatalf("run() returned %v before every worker exited", got)
	default:
	}

	close(release)
	select {
	case got := <-returned:
		sort.Ints(got)
		if want := []int{1, 4, 9}; !reflect.DeepEqual(got, want) {
			t.Fatalf("run() = %v, want %v", got, want)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("run() did not return after every worker exited")
	}
}
