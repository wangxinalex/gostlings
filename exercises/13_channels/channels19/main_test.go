package main

import (
	"testing"
	"time"
)

func TestRunCancelsAndJoinsTheSlowProducer(t *testing.T) {
	done := make(chan struct{})
	returned := make(chan string, 1)
	go func() { returned <- run(done) }()

	select {
	case got := <-returned:
		if got != "timed out" {
			t.Fatalf("run() = %q, want %q", got, "timed out")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("run() did not time out and cancel its producer")
	}

	select {
	case <-done:
	default:
		t.Fatal("run() returned before its producer finished")
	}
}
