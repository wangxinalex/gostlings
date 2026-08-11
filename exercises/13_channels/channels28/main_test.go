package main

import (
	"errors"
	"testing"
	"time"
)

func TestRunReturnsTheFirstObservedError(t *testing.T) {
	want := errors.New("bad job")
	if got := run(2, []job{{value: 1, err: want}, {value: 2}}); !errors.Is(got, want) {
		t.Fatalf("run() error = %v, want %v", got, want)
	}
}

func TestRunInspectsAnErrorAfterASuccessfulJob(t *testing.T) {
	want := errors.New("later bad job")
	if got := run(1, []job{{value: 1}, {value: 2, err: want}}); !errors.Is(got, want) {
		t.Fatalf("run() error = %v, want %v", got, want)
	}
}

func TestRunClosesStopOnceAndJoinsWorkersBeforeReturning(t *testing.T) {
	want := errors.New("bad job")
	previousStop := onStopClosed
	previousExit := onWorkerExit
	stopped := make(chan struct{}, 1)
	exited := make(chan struct{}, 2)
	release := make(chan struct{})
	onStopClosed = func() { stopped <- struct{}{} }
	onWorkerExit = func() {
		exited <- struct{}{}
		<-release
	}
	t.Cleanup(func() {
		onStopClosed = previousStop
		onWorkerExit = previousExit
	})

	returned := make(chan error, 1)
	go func() { returned <- run(2, []job{{value: 1, err: want}, {value: 2}}) }()
	select {
	case <-stopped:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("run() did not close stop after an error")
	}
	for worker := 0; worker < 2; worker++ {
		select {
		case <-exited:
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("worker %d did not observe stop and exit", worker+1)
		}
	}
	select {
	case got := <-returned:
		t.Fatalf("run() returned %v before all workers joined", got)
	default:
	}

	close(release)
	select {
	case got := <-returned:
		if !errors.Is(got, want) {
			t.Fatalf("run() error = %v, want %v", got, want)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("run() did not return after workers joined")
	}
	select {
	case <-stopped:
		t.Fatal("run() closed stop more than once")
	default:
	}
}
