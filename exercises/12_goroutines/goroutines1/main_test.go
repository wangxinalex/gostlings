package main

import (
	"fmt"
	"gostlings/internal/testutil"
	"runtime"
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
		runtime.Goexit()
	}
	t.Cleanup(func() { greeting = originalGreeting })

	output := make(chan string, 1)
	returned := make(chan struct{}, 1)
	go func() {
		output <- testutil.CaptureStdout(t, func() {
			main()
			returned <- struct{}{}
		})
	}()
	select {
	case <-started:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("greeting task did not start")
	}

	close(release)
	select {
	case <-returned:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("main did not return after greeting was released")
	}
	select {
	case got := <-output:
		if got != "hello\n" {
			t.Fatalf("main output = %q, want %q", got, "hello\n")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("greeting output was not captured")
	}
}
