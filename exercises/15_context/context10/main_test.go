package main

import (
	"context"
	"testing"
	"time"
)

func TestRunWorkersJoinsEveryCanceledWorker(t *testing.T) {
	const count = 4
	previousStarted, previousStopped := workerStarted, workerStopped
	started := make(chan struct{}, count)
	stopped := make(chan struct{}, count)
	workerStarted = func() { started <- struct{}{} }
	workerStopped = func() { stopped <- struct{}{} }
	t.Cleanup(func() { workerStarted, workerStopped = previousStarted, previousStopped })

	ctx, cancel := context.WithCancel(context.Background())
	done := runWorkers(ctx, count)
	for i := 0; i < count; i++ {
		select {
		case <-started:
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("worker %d did not start", i)
		}
	}
	cancel()
	for i := 0; i < count; i++ {
		select {
		case <-stopped:
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("worker %d did not stop after cancellation", i)
		}
	}
	select {
	case <-done:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("done did not close after every worker stopped")
	}
}

func TestRunWorkersWithNoWorkersClosesDone(t *testing.T) {
	select {
	case <-runWorkers(context.Background(), 0):
	case <-time.After(500 * time.Millisecond):
		t.Fatal("done did not close for zero workers")
	}
}
