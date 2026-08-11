package main

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestStartPoolExposesDirectionalStreamsAndClosesResults(t *testing.T) {
	var jobs chan<- int
	var results <-chan int
	jobs, results = startPool(2)
	go func() {
		for _, job := range []int{1, 2, 3, 4} {
			jobs <- job
		}
		close(jobs)
	}()

	got := collectPoolResults(t, results)
	sort.Ints(got)
	if want := []int{1, 4, 9, 16}; !reflect.DeepEqual(got, want) {
		t.Fatalf("startPool() results = %v, want %v", got, want)
	}
}

func TestStartPoolStartsEveryRequestedWorker(t *testing.T) {
	previous := onPoolWorkerStart
	started := make(chan struct{}, 3)
	onPoolWorkerStart = func() { started <- struct{}{} }
	t.Cleanup(func() { onPoolWorkerStart = previous })

	jobs, results := startPool(3)
	for worker := 0; worker < 3; worker++ {
		select {
		case <-started:
		case <-time.After(500 * time.Millisecond):
			t.Fatal("startPool() did not start every requested worker")
		}
	}
	close(jobs)
	collectPoolResults(t, results)
}

func TestStartPoolClosesResultsAfterEveryWorkerExits(t *testing.T) {
	previous := onPoolWorkerExit
	exited := make(chan struct{}, 2)
	release := make(chan struct{})
	onPoolWorkerExit = func() {
		exited <- struct{}{}
		<-release
	}
	t.Cleanup(func() { onPoolWorkerExit = previous })

	jobs, results := startPool(2)
	close(jobs)
	for worker := 0; worker < 2; worker++ {
		select {
		case <-exited:
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("worker %d did not exit", worker+1)
		}
	}
	select {
	case _, ok := <-results:
		if !ok {
			t.Fatal("startPool() closed results before every worker exited")
		}
		t.Fatal("startPool() emitted a result without a job")
	default:
	}
	close(release)
	collectPoolResults(t, results)
}

func collectPoolResults(t *testing.T, results <-chan int) []int {
	t.Helper()
	var got []int
	for {
		select {
		case value, ok := <-results:
			if !ok {
				return got
			}
			got = append(got, value)
		case <-time.After(500 * time.Millisecond):
			t.Fatal("startPool() results did not close")
		}
	}
}
