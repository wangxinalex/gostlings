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

func waitForWorkersDone(t *testing.T, done <-chan struct{}) {
	t.Helper()
	select {
	case <-done:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("startWorkers() did not finish")
	}
}
