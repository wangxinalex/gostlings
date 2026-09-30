package main

import (
	"errors"
	"testing"
	"time"
)

func TestRunAllReturnsNilWhenEveryJobSucceeds(t *testing.T) {
	done := make(chan struct{}, 3)
	jobs := []func() error{
		func() error { done <- struct{}{}; return nil },
		func() error { done <- struct{}{}; return nil },
		func() error { done <- struct{}{}; return nil },
	}
	if err := runAll(jobs); err != nil {
		t.Fatalf("runAll() = %v, want nil", err)
	}
	if len(done) != 3 {
		t.Fatalf("%d jobs completed, want 3", len(done))
	}
}

func TestRunAllReturnsTheFirstFailure(t *testing.T) {
	first := errors.New("first failure")
	later := errors.New("later failure")
	jobs := []func() error{
		func() error { return nil },
		func() error { return first },
		func() error { return later },
	}
	if err := runAll(jobs); !errors.Is(err, first) {
		t.Fatalf("runAll() = %v, want %v", err, first)
	}
}

func TestRunAllWaitsForEveryJobAfterAFailure(t *testing.T) {
	wantErr := errors.New("boom")
	started := make(chan struct{})
	release := make(chan struct{})
	finished := make(chan struct{}, 1)

	result := make(chan error, 1)
	go func() {
		result <- runAll([]func() error{
			func() error {
				close(started)
				<-release
				finished <- struct{}{}
				return nil
			},
			func() error { return wantErr },
		})
	}()

	select {
	case <-started:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("the gated job did not start")
	}

	// A failing job must not make runAll return while another job is still running.
	select {
	case got := <-result:
		t.Fatalf("runAll returned %v before every job finished", got)
	default:
	}

	close(release)
	select {
	case got := <-result:
		if !errors.Is(got, wantErr) {
			t.Fatalf("runAll() = %v, want %v", got, wantErr)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runAll did not return after every job finished")
	}

	select {
	case <-finished:
	default:
		t.Fatal("the gated job did not run to completion")
	}
}

func TestRunAllHandlesNoJobs(t *testing.T) {
	if err := runAll(nil); err != nil {
		t.Fatalf("runAll(nil) = %v, want nil", err)
	}
	if err := runAll([]func() error{}); err != nil {
		t.Fatalf("runAll([]) = %v, want nil", err)
	}
}
