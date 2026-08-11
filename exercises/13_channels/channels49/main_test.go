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
		_, err := runFirstErrorBounded(make(chan struct{}), 2, []job{{value: 2}, {err: bad}})
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

func TestRunFirstErrorBoundedExternalStopWaitsForActiveWork(t *testing.T) {
	previous := processFirstErrorBounded
	started := make(chan struct{}, 1)
	release := make(chan struct{})
	processFirstErrorBounded = func(value int) int {
		started <- struct{}{}
		<-release
		return value * value
	}
	t.Cleanup(func() { processFirstErrorBounded = previous })

	stop := make(chan struct{})
	returned := make(chan struct {
		values []int
		err    error
	}, 1)
	go func() {
		values, err := runFirstErrorBounded(stop, 1, []job{{value: 2}})
		returned <- struct {
			values []int
			err    error
		}{values, err}
	}()
	select {
	case <-started:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runFirstErrorBounded() did not start active work")
	}
	close(stop)
	select {
	case got := <-returned:
		t.Fatalf("runFirstErrorBounded() returned (%v, %v) before active work exited", got.values, got.err)
	default:
	}
	close(release)
	select {
	case got := <-returned:
		if got.err != nil || len(got.values) != 0 {
			t.Fatalf("runFirstErrorBounded() = (%v, %v), want no successful values and no job error after external stop", got.values, got.err)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runFirstErrorBounded() did not join active work after external stop")
	}
}

func TestRunFirstErrorBoundedLimitsActiveSuccessfulWork(t *testing.T) {
	previous := processFirstErrorBounded
	started := make(chan int, 2)
	activeSlots := make(chan struct{}, 2)
	activeSlots <- struct{}{}
	activeSlots <- struct{}{}
	overflow := make(chan int, 1)
	release := make(chan struct{})
	processFirstErrorBounded = func(value int) int {
		select {
		case <-activeSlots:
			started <- value
		default:
			select {
			case <-release:
			default:
				overflow <- value
			}
		}
		<-release
		return value * value
	}
	t.Cleanup(func() { processFirstErrorBounded = previous })

	returned := make(chan struct {
		values []int
		err    error
	}, 1)
	go func() {
		values, err := runFirstErrorBounded(make(chan struct{}), 2, []job{{value: 3}, {value: 1}, {value: 2}})
		returned <- struct {
			values []int
			err    error
		}{values, err}
	}()
	for worker := 0; worker < 2; worker++ {
		select {
		case <-started:
		case <-time.After(500 * time.Millisecond):
			t.Fatal("runFirstErrorBounded() did not start work up to its worker limit")
		}
	}
	close(release)
	select {
	case got := <-returned:
		sort.Ints(got.values)
		if got.err != nil || !reflect.DeepEqual(got.values, []int{1, 4, 9}) {
			t.Fatalf("runFirstErrorBounded() = (%v, %v), want all successful values", got.values, got.err)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runFirstErrorBounded() did not finish after active work was released")
	}
	select {
	case value := <-overflow:
		t.Fatalf("runFirstErrorBounded() started a third active job (%d) with workers 2", value)
	default:
	}
}
