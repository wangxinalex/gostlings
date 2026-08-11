package main

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestSquareWorkersForwardsEverySquareAndCloses(t *testing.T) {
	jobs := make(chan int, 4)
	for _, job := range []int{1, 2, 3, 4} {
		jobs <- job
	}
	close(jobs)

	got := collectCancellableSquares(t, squareWorkers(make(chan struct{}), 2, jobs))
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
	out := squareWorkers(make(chan struct{}), 3, jobs)
	for worker := 0; worker < 3; worker++ {
		select {
		case <-started:
		case <-time.After(500 * time.Millisecond):
			t.Fatal("squareWorkers() did not start every requested worker")
		}
	}
	close(jobs)
	collectCancellableSquares(t, out)
}

func TestSquareWorkersStopsWhileWaitingForJobs(t *testing.T) {
	stop := make(chan struct{})
	jobs := make(chan int)
	out := squareWorkers(stop, 2, jobs)
	close(stop)
	waitForCancellableSquaresClose(t, out)
}

func TestSquareWorkersStopsWhileAResultSendIsBlocked(t *testing.T) {
	previous := onSquareWorkerBeforeSend
	beforeSend := make(chan struct{}, 1)
	onSquareWorkerBeforeSend = func() { beforeSend <- struct{}{} }
	t.Cleanup(func() { onSquareWorkerBeforeSend = previous })

	stop := make(chan struct{})
	jobs := make(chan int)
	out := squareWorkers(stop, 1, jobs)
	sent := make(chan struct{})
	go func() {
		jobs <- 5
		close(sent)
	}()
	select {
	case <-sent:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("squareWorkers() did not receive its job")
	}
	select {
	case <-beforeSend:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("squareWorkers() did not begin its result send")
	}

	close(stop)
	waitForCancellableSquaresClose(t, out)
}

func collectCancellableSquares(t *testing.T, out <-chan int) []int {
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
			t.Fatal("squareWorkers() did not close")
		}
	}
}

func waitForCancellableSquaresClose(t *testing.T, out <-chan int) {
	t.Helper()
	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("squareWorkers() sent a value after cancellation")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("squareWorkers() did not close after cancellation")
	}
}
