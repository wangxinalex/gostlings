package main

import (
	"reflect"
	"testing"
	"time"
)

func TestRunOrderedBoundedRestoresOrderAndHandlesEmptyJobs(t *testing.T) {
	got, ok := runOrderedBounded(make(chan struct{}), 2, []int{3, 1, 2})
	if !ok || !reflect.DeepEqual(got, []int{9, 1, 4}) {
		t.Fatalf("runOrderedBounded() = (%v, %v), want ([9 1 4], true)", got, ok)
	}
	got, ok = runOrderedBounded(make(chan struct{}), 2, nil)
	if !ok || len(got) != 0 {
		t.Fatalf("runOrderedBounded(nil) = (%v, %v), want ([], true)", got, ok)
	}
}

func TestRunOrderedBoundedLimitsWorkersAndCancelsBeforeLaterJobsStart(t *testing.T) {
	previous := processOrderedBounded
	started := make(chan int, 2)
	activeSlots := make(chan struct{}, 2)
	activeSlots <- struct{}{}
	activeSlots <- struct{}{}
	overflow := make(chan int, 1)
	release := make(chan struct{})
	processOrderedBounded = func(value int) int {
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
	t.Cleanup(func() { processOrderedBounded = previous })

	stop := make(chan struct{})
	returned := make(chan struct {
		values []int
		ok     bool
	}, 1)
	go func() {
		values, ok := runOrderedBounded(stop, 2, []int{1, 2, 3})
		returned <- struct {
			values []int
			ok     bool
		}{values, ok}
	}()
	for worker := 0; worker < 2; worker++ {
		select {
		case <-started:
		case <-time.After(500 * time.Millisecond):
			t.Fatal("runOrderedBounded() did not start its bounded workers")
		}
	}
	close(stop)
	close(release)
	select {
	case got := <-returned:
		if got.ok {
			t.Fatalf("runOrderedBounded() = (%v, true), want cancellation", got.values)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runOrderedBounded() did not join canceled workers")
	}
	select {
	case value := <-overflow:
		t.Fatalf("runOrderedBounded() started later job %d before a worker was released", value)
	default:
	}
}
