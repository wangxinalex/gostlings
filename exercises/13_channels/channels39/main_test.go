package main

import (
	"testing"
	"time"
)

func TestServeCommandsPausesThenResumesJobs(t *testing.T) {
	previous := onCommandApplied
	applied := make(chan command, 2)
	pauseApplied := make(chan struct{})
	onCommandApplied = func(c command) {
		if c == pause {
			close(pauseApplied)
		}
		applied <- c
	}
	t.Cleanup(func() { onCommandApplied = previous })
	stop, commands, jobs := make(chan struct{}), make(chan command, 2), make(chan int)
	out := serveCommands(stop, commands, jobs)
	commands <- pause
	wait39Signal(t, pauseApplied, "serveCommands() did not begin applying pause")
	sent := make(chan struct{})
	ready := make(chan struct{})
	go func() { close(ready); jobs <- 7; close(sent) }()
	wait39Signal(t, ready, "job sender did not become ready")
	select {
	case <-sent:
		t.Fatal("serveCommands() consumed a job while paused")
	case <-time.After(50 * time.Millisecond):
	}
	commands <- resume
	wait39(t, applied, "serveCommands() did not apply resume")
	select {
	case <-sent:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("serveCommands() did not receive the resumed job")
	}
	select {
	case got := <-out:
		if got != 7 {
			t.Fatalf("serveCommands() = %d, want 7", got)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("serveCommands() did not forward resumed job")
	}
	close(jobs)
	closed39(t, out)
}
func TestServeCommandsStopsWhilePausedAndJobsAreBlocked(t *testing.T) {
	previous := onCommandApplied
	applied := make(chan command, 1)
	onCommandApplied = func(c command) { applied <- c }
	t.Cleanup(func() { onCommandApplied = previous })
	stop, commands, jobs := make(chan struct{}), make(chan command, 1), make(chan int)
	out := serveCommands(stop, commands, jobs)
	commands <- pause
	wait39(t, applied, "serveCommands() did not apply pause")
	close(stop)
	closed39(t, out)
}

func TestServeCommandsStopsWhenCommandInputCloses(t *testing.T) {
	commands := make(chan command)
	out := serveCommands(make(chan struct{}), commands, make(chan int))
	close(commands)
	closed39(t, out)
}

func wait39(t *testing.T, ch <-chan command, message string) {
	t.Helper()
	select {
	case <-ch:
	case <-time.After(500 * time.Millisecond):
		t.Fatal(message)
	}
}

func wait39Signal(t *testing.T, ch <-chan struct{}, message string) {
	t.Helper()
	select {
	case <-ch:
	case <-time.After(500 * time.Millisecond):
		t.Fatal(message)
	}
}

func closed39(t *testing.T, out <-chan int) {
	t.Helper()
	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("serveCommands() sent after shutdown")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("serveCommands() did not close after shutdown")
	}
}
