package main

import (
	"sync"
	"testing"
	"time"
)

func TestUpdateStateCommitsTheTransformedSnapshot(t *testing.T) {
	state := &protectedState{value: 10}
	if got := state.updateState(func(value int) int { return value + 5 }); got != 15 {
		t.Fatalf("updateState() = %d, want 15", got)
	}
	if got := state.valueSnapshot(); got != 15 {
		t.Fatalf("valueSnapshot() = %d, want 15", got)
	}
}

func TestUpdateStateDoesNotHoldTheLockDuringSlowWork(t *testing.T) {
	state := &protectedState{value: 1}
	entered := make(chan struct{})
	release := make(chan struct{})
	finished := make(chan struct{})
	go func() {
		state.updateState(func(value int) int {
			close(entered)
			<-release
			return value + 1
		})
		close(finished)
	}()
	select {
	case <-entered:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("transform was not called")
	}
	readDone := make(chan struct{})
	go func() {
		_ = state.valueSnapshot()
		close(readDone)
	}()
	select {
	case <-readDone:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("valueSnapshot() blocked while transform was running; lock was held too long")
	}
	close(release)
	select {
	case <-finished:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("updateState() did not finish")
	}
}

func TestUpdateStateIsSafeForConcurrentUpdates(t *testing.T) {
	state := &protectedState{}
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			state.updateState(func(value int) int { return value + 1 })
		}()
	}
	wg.Wait()
	if got := state.valueSnapshot(); got < 1 || got > 20 {
		t.Fatalf("valueSnapshot() = %d, want a committed update", got)
	}
}
