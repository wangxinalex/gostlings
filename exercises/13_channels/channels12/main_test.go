package main

import (
	"testing"
	"time"
)

func TestStartWorkersStopsEveryWorker(t *testing.T) {
	stop := make(chan struct{})
	done := startWorkers(3, stop)
	close(stop)

	select {
	case <-done:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("not all workers stopped after broadcast cancellation")
	}
}

func TestStartWorkersWithNoWorkersCompletes(t *testing.T) {
	done := startWorkers(0, make(chan struct{}))
	select {
	case <-done:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("startWorkers(0) did not close done")
	}
}
