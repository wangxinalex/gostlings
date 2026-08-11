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

func TestRunJoinsTheSuccessfulProducerBeforeReturning(t *testing.T) {
	previousProducer := runProducer
	runProducer = func(stop <-chan struct{}, result chan<- string) {
		select {
		case result <- "finished":
		case <-stop:
		}
	}
	t.Cleanup(func() { runProducer = previousProducer })

	done := make(chan struct{})
	if got := run(done); got != "finished" {
		t.Fatalf("run() = %q, want %q", got, "finished")
	}
	select {
	case <-done:
	default:
		t.Fatal("run() returned before its successful producer finished")
	}
}
