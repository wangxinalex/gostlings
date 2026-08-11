package main

import (
	"gostlings/internal/testutil"
	"testing"
	"time"
)

func TestMainCompletesGreeting(t *testing.T) {
	output := make(chan string, 1)
	go func() {
		output <- testutil.CaptureStdout(t, main)
	}()

	select {
	case got := <-output:
		if got != "hello\n" {
			t.Fatalf("main output = %q, want %q", got, "hello\n")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("main did not finish its greeting")
	}
}
