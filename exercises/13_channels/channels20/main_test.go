package main

import (
	"testing"
	"time"
)

func TestShutdownClosesStopThenReportsAllWorkersDone(t *testing.T) {
	stop := make(chan struct{})
	done := shutdown(stop, 3)

	select {
	case <-stop:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("shutdown() did not close stop")
	}

	waitForShutdownDone(t, done)
}

func TestShutdownDoesNotCloseAnAlreadyClosedStopAgain(t *testing.T) {
	stop := make(chan struct{})
	first := shutdown(stop, 2)
	waitForShutdownDone(t, first)

	second := shutdown(stop, 1)
	waitForShutdownDone(t, second)
}

func TestShutdownWaitsForEveryWorkerExit(t *testing.T) {
	const workers = 3
	exited := make(chan struct{})
	previousHook := onShutdownWorkerExit
	onShutdownWorkerExit = func() { exited <- struct{}{} }
	t.Cleanup(func() { onShutdownWorkerExit = previousHook })

	stop := make(chan struct{})
	done := shutdown(stop, workers)

	for worker := 0; worker < workers-1; worker++ {
		select {
		case <-exited:
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("worker %d did not exit", worker+1)
		}
		select {
		case <-done:
			t.Fatal("shutdown() completed before every worker exited")
		default:
		}
	}

	select {
	case <-exited:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("last worker did not exit")
	}
	waitForShutdownDone(t, done)
}

func waitForShutdownDone(t *testing.T, done <-chan struct{}) {
	t.Helper()
	select {
	case <-done:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("shutdown() did not finish")
	}
}
