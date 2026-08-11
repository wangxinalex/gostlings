package main

import (
	"reflect"
	"testing"
	"time"
)

func TestParallelReturnsOrderedResultsWhenNotStopped(t *testing.T) {
	got, ok := parallel(make(chan struct{}), 2, []int{3, 1, 2}, func(value int) int { return value * value })
	if !ok || !reflect.DeepEqual(got, []int{9, 1, 4}) {
		t.Fatalf("parallel() = (%v, %v), want ([9 1 4], true)", got, ok)
	}
}

func TestParallelStopsTokenAcquisitionAndBlockedResultPublication(t *testing.T) {
	stop := make(chan struct{})
	started := make(chan struct{}, 1)
	release := make(chan struct{})
	returned := make(chan struct {
		values []int
		ok     bool
	}, 1)
	go func() {
		values, ok := parallel(stop, 1, []int{1, 2}, func(value int) int {
			started <- struct{}{}
			<-release
			return value
		})
		returned <- struct {
			values []int
			ok     bool
		}{values, ok}
	}()

	select {
	case <-started:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("parallel() did not start its first job")
	}
	close(stop)
	close(release)
	select {
	case got := <-returned:
		if got.ok {
			t.Fatalf("parallel() = (%v, true), want cancellation", got.values)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("parallel() did not return after cancellation released blocked work")
	}
}
