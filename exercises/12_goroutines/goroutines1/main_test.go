package main

import (
	"fmt"
	"gostlings/internal/testutil"
	"testing"
	"time"
)

func TestMainLaunchesAndJoinsGreeting(t *testing.T) {
	originalGreeting := greeting
	started := make(chan struct{})
	release := make(chan struct{})
	greeting = func() {
		close(started)
		<-release
		fmt.Println("hello")
	}
	t.Cleanup(func() { greeting = originalGreeting })

	output := make(chan string, 1)
	go func() { output <- testutil.CaptureStdout(t, main) }()
	select {
	case <-started:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("greeting task did not start")
	}
	select {
	case got := <-output:
		t.Fatalf("main returned while greeting was blocked; output = %q", got)
	case <-time.After(100 * time.Millisecond):
	}

	close(release)
	select {
	case got := <-output:
		if got != "hello\n" {
			t.Fatalf("main output = %q, want %q", got, "hello\n")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("main did not return after greeting was released")
	}
}
