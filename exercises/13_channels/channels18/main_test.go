package main

import (
	"testing"
	"time"
)

func TestStartWorkersWaitsForBroadcastCancellation(t *testing.T) {
	stop := make(chan struct{})
	done := startWorkers(3, stop)
	select {
	case <-done:
		t.Fatal("startWorkers() completed before cancellation")
	default:
	}

	close(stop)
	waitForWorkersDone(t, done)
}

func TestStartWorkersWithNoWorkersCompletes(t *testing.T) {
	waitForWorkersDone(t, startWorkers(0, make(chan struct{})))
}

func TestStartWorkersWaitsForEveryWorkerExit(t *testing.T) {
	const workers = 3
	exited := make(chan struct{})
	previousHook := onWorkerExit
	onWorkerExit = func() { exited <- struct{}{} }
	t.Cleanup(func() { onWorkerExit = previousHook })

	stop := make(chan struct{})
	done := startWorkers(workers, stop)
	close(stop)

	for worker := 0; worker < workers-1; worker++ {
		select {
		case <-exited:
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("worker %d did not exit", worker+1)
		}
		select {
		case <-done:
			t.Fatal("startWorkers() completed before every worker exited")
		default:
		}
	}

	select {
	case <-exited:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("last worker did not exit")
	}
	waitForWorkersDone(t, done)
}

func waitForWorkersDone(t *testing.T, done <-chan struct{}) {
	t.Helper()
	select {
	case <-done:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("startWorkers() did not finish")
	}
}
