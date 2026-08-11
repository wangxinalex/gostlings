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
	greeting = func() {
		close(started)
		fmt.Println("hello")
		runtime.Goexit()
	}
	t.Cleanup(func() { greeting = originalGreeting })

	got := testutil.CaptureStdout(t, func() {
		returned := make(chan struct{}, 1)
		go func() {
			main()
			returned <- struct{}{}
		}()

		select {
		case <-returned:
		case <-time.After(500 * time.Millisecond):
			t.Fatal("main did not return after the greeting worker exited")
		}
	})
	if got != "hello\n" {
		t.Fatalf("main output = %q, want %q", got, "hello\n")
	}
	select {
	case <-started:
	default:
		t.Fatal("main returned without launching the greeting task")
	}
}
