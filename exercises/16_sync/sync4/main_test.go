package main

import (
	"sync"
	"testing"
	"time"
)

func TestWaitForWorkersRunsAndJoinsEveryWorker(t *testing.T) {
	const count = 5
	previous := workerDone
	started := make(chan struct{}, count)
	workerDone = func() { started <- struct{}{} }
	t.Cleanup(func() { workerDone = previous })
	if got := waitForWorkers(count); got != count {
		t.Fatalf("waitForWorkers() = %d, want %d", got, count)
	}
	for i := 0; i < count; i++ {
		select {
		case <-started:
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("worker %d was not started", i)
		}
	}
}

func TestWaitForWorkersHandlesEmptyInput(t *testing.T) {
	var guard sync.Mutex
	called := 0
	previous := workerDone
	workerDone = func() { guard.Lock(); called++; guard.Unlock() }
	t.Cleanup(func() { workerDone = previous })
	if got := waitForWorkers(0); got != 0 || called != 0 {
		t.Fatalf("waitForWorkers(0) = %d with %d worker calls, want 0 and 0", got, called)
	}
}
