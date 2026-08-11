package main

import (
	"errors"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestRunFirstErrorBoundedReturnsSuccessfulValues(t *testing.T) {
	got, err := runFirstErrorBounded(make(chan struct{}), 2, []job{{value: 3}, {value: 1}, {value: 2}})
	sort.Ints(got)
	if err != nil || !reflect.DeepEqual(got, []int{1, 4, 9}) {
		t.Fatalf("runFirstErrorBounded() = (%v, %v), want ([1 4 9], nil)", got, err)
	}
}

func TestRunFirstErrorBoundedStopsLaterWorkAndReturnsTheJobError(t *testing.T) {
	previous := processFirstErrorBounded
	processed := make(chan int, 1)
	processFirstErrorBounded = func(value int) int { processed <- value; return value * value }
	t.Cleanup(func() { processFirstErrorBounded = previous })

	bad := errors.New("bad job")
	returned := make(chan struct {
		values []int
		err    error
	}, 1)
	go func() {
		values, err := runFirstErrorBounded(make(chan struct{}), 1, []job{{value: 1, err: bad}, {value: 2}})
		returned <- struct {
			values []int
			err    error
		}{values, err}
	}()
	select {
	case got := <-returned:
		if !errors.Is(got.err, bad) {
			t.Fatalf("runFirstErrorBounded() error = %v, want %v", got.err, bad)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runFirstErrorBounded() did not return the first error")
	}
	select {
	case value := <-processed:
		t.Fatalf("runFirstErrorBounded() processed later value %d after an error", value)
	default:
	}
}

func TestRunFirstErrorBoundedWaitsForActiveWorkerBeforeReturningError(t *testing.T) {
	previous := processFirstErrorBounded
	started := make(chan struct{}, 1)
	release := make(chan struct{})
	processFirstErrorBounded = func(value int) int {
		started <- struct{}{}
		<-release
		return value * value
	}
	t.Cleanup(func() { processFirstErrorBounded = previous })

	bad := errors.New("bad job")
	returned := make(chan error, 1)
	go func() {
		_, err := runFirstErrorBounded(make(chan struct{}), 2, []job{{err: bad}, {value: 2}})
		returned <- err
	}()
	select {
	case <-started:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runFirstErrorBounded() did not start work alongside the failing job")
	}
	select {
	case err := <-returned:
		t.Fatalf("runFirstErrorBounded() returned %v before its active worker exited", err)
	default:
	}
	close(release)
	select {
	case err := <-returned:
		if !errors.Is(err, bad) {
			t.Fatalf("runFirstErrorBounded() error = %v, want %v", err, bad)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runFirstErrorBounded() did not return after the active worker exited")
	}
}
