package main

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestSquareWorkersProcessesEveryJobAndClosesOutput(t *testing.T) {
	jobs := make(chan int, 4)
	for _, job := range []int{1, 2, 3, 4} {
		jobs <- job
	}
	close(jobs)

	got := collectSquares(t, squareWorkers(2, jobs))
	sort.Ints(got)
	if want := []int{1, 4, 9, 16}; !reflect.DeepEqual(got, want) {
		t.Fatalf("squareWorkers() = %v, want %v", got, want)
	}
}

func TestSquareWorkersStartsEachRequestedWorker(t *testing.T) {
	previous := onSquareWorkerStart
	started := make(chan struct{}, 3)
	onSquareWorkerStart = func() { started <- struct{}{} }
	t.Cleanup(func() { onSquareWorkerStart = previous })

	jobs := make(chan int)
	out := squareWorkers(3, jobs)
	for worker := 0; worker < 3; worker++ {
		waitForSquareWorkerStart(t, started)
	}
	close(jobs)
	collectSquares(t, out)
}

func TestSquareWorkersClosesOnlyAfterEveryWorkerExits(t *testing.T) {
	previous := onSquareWorkerExit
	exited := make(chan struct{}, 2)
	release := make(chan struct{})
	onSquareWorkerExit = func() {
		exited <- struct{}{}
		<-release
	}
	t.Cleanup(func() { onSquareWorkerExit = previous })

	jobs := make(chan int)
	close(jobs)
	out := squareWorkers(2, jobs)
	for worker := 0; worker < 2; worker++ {
		select {
		case <-exited:
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("worker %d did not exit", worker+1)
		}
	}
	select {
	case _, ok := <-out:
		if !ok {
			t.Fatal("squareWorkers() closed output before every worker exited")
		}
		t.Fatal("squareWorkers() emitted a result without a job")
	default:
	}

	close(release)
	collectSquares(t, out)
}

func collectSquares(t *testing.T, out <-chan int) []int {
	t.Helper()
	var got []int
	for {
		select {
		case value, ok := <-out:
			if !ok {
				return got
			}
			got = append(got, value)
		case <-time.After(500 * time.Millisecond):
			t.Fatal("output did not close")
		}
	}
}

func waitForSquareWorkerStart(t *testing.T, started <-chan struct{}) {
	t.Helper()
	select {
	case <-started:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("squareWorkers() did not start every requested worker")
	}
}
