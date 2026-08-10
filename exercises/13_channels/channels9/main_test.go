package main

import (
	"testing"
	"time"
)

func TestCompleteClosesDoneChannel(t *testing.T) {
	done := complete()
	select {
	case <-done:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("complete() did not close its done channel")
	}
}
