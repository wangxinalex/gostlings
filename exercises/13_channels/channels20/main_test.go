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

func waitForShutdownDone(t *testing.T, done <-chan struct{}) {
	t.Helper()
	select {
	case <-done:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("shutdown() did not finish")
	}
}
