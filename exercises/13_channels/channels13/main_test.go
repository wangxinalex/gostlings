package main

import (
	"testing"
	"time"
)

func TestCompleteClosesADoneChannelWithoutAValue(t *testing.T) {
	done := complete()
	select {
	case _, ok := <-done:
		if ok {
			t.Fatal("complete() sent data on its done channel")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("complete() did not close its done channel")
	}
}
