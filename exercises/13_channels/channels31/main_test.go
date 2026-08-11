package main

import (
	"testing"
	"time"
)

func TestFirstResultCancelsRemainingTasksBeforeReturning(t *testing.T) {
	started := make(chan struct{}, 3)
	releaseWinner := make(chan struct{})
	releaseCanceled := make(chan struct{})
	cancelled := make(chan struct{}, 2)
	tasks := []func(<-chan struct{}) string{
		func(stop <-chan struct{}) string { started <- struct{}{}; <-releaseWinner; return "winner" },
		func(stop <-chan struct{}) string {
			started <- struct{}{}
			<-stop
			cancelled <- struct{}{}
			<-releaseCanceled
			return "late"
		},
		func(stop <-chan struct{}) string {
			started <- struct{}{}
			<-stop
			cancelled <- struct{}{}
			return "late"
		},
	}
	returned := make(chan string, 1)
	go func() { returned <- firstResult(tasks) }()
	for range tasks {
		wait31(t, started, "firstResult() did not start every task")
	}
	close(releaseWinner)
	for range tasks[1:] {
		wait31(t, cancelled, "firstResult() returned before cancellation reached every remaining task")
	}
	select {
	case got := <-returned:
		t.Fatalf("firstResult() returned %q before canceled tasks exited", got)
	default:
	}
	close(releaseCanceled)
	select {
	case got := <-returned:
		if got != "winner" {
			t.Fatalf("firstResult() = %q, want %q", got, "winner")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("firstResult() did not return after canceled tasks exited")
	}
}

func TestFirstResultWithNoTasksReturnsEmptyString(t *testing.T) {
	if got := firstResult(nil); got != "" {
		t.Fatalf("firstResult(nil) = %q, want empty string", got)
	}
}

func wait31(t *testing.T, ch <-chan struct{}, message string) {
	t.Helper()
	select {
	case <-ch:
	case <-time.After(500 * time.Millisecond):
		t.Fatal(message)
	}
}
