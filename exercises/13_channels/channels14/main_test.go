package main

import (
	"testing"
	"time"
)

func TestRunAsyncPublishesBeforeTheCallerReceives(t *testing.T) {
	workStarted := make(chan struct{})
	returned := make(chan (<-chan int), 1)
	go func() {
		returned <- runAsync(func() int {
			close(workStarted)
			return 42
		})
	}()

	select {
	case <-workStarted:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runAsync() did not start work")
	}

	var result <-chan int
	select {
	case result = <-returned:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runAsync() did not return after work published its result")
	}
	if got := cap(result); got != 1 {
		t.Fatalf("cap(runAsync() result) = %d, want 1", got)
	}

	select {
	case got := <-result:
		if got != 42 {
			t.Fatalf("runAsync() result = %d, want 42", got)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runAsync() did not publish its result")
	}
}
