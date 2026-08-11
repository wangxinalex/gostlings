package main

import (
	"testing"
	"time"
)

func TestWaitTimerDeliversOneEvent(t *testing.T) {
	stop := make(chan struct{})
	result := waitTimer(stop, 10*time.Millisecond)
	select {
	case _, ok := <-result:
		if !ok {
			t.Fatal("waitTimer() closed before its timer fired")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("waitTimer() did not deliver its timer event")
	}
	select {
	case _, ok := <-result:
		if ok {
			t.Fatal("waitTimer() delivered more than one event")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("waitTimer() did not close after its one event")
	}
}

func TestWaitTimerClosesWhenStoppedEarly(t *testing.T) {
	stop := make(chan struct{})
	result := waitTimer(stop, time.Hour)
	close(stop)
	select {
	case _, ok := <-result:
		if ok {
			t.Fatal("waitTimer() delivered an event after early stop")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("waitTimer() did not close after early stop")
	}
}
