package main

import (
	"testing"
	"time"
)

func TestBoundedStateBlocksAtLimitAndReleasesCapacity(t *testing.T) {
	state := newBoundedState(1)
	if !state.Acquire() {
		t.Fatal("first Acquire() failed")
	}
	acquired := make(chan bool, 1)
	go func() { acquired <- state.Acquire() }()
	select {
	case <-acquired:
		t.Fatal("second Acquire() passed the limit")
	case <-time.After(20 * time.Millisecond):
	}
	state.Release()
	select {
	case ok := <-acquired:
		if !ok || state.Active() != 1 {
			t.Fatalf("second Acquire() = %v with active %d, want true and 1", ok, state.Active())
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("Release() did not wake a waiting Acquire")
	}
	state.Release()
}

func TestBoundedStateCloseRejectsAndWakesWaiters(t *testing.T) {
	state := newBoundedState(1)
	state.Acquire()
	acquired := make(chan bool, 1)
	go func() { acquired <- state.Acquire() }()
	select {
	case <-acquired:
		t.Fatal("waiter acquired before close")
	case <-time.After(20 * time.Millisecond):
	}
	state.Close()
	select {
	case ok := <-acquired:
		if ok {
			t.Fatal("Acquire() succeeded after Close")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("Close() did not wake blocked Acquire")
	}
	state.Release()
	if state.Acquire() {
		t.Fatal("Acquire() succeeded after Close")
	}
}
