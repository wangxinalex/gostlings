package main

import (
	"testing"
	"time"
)

func TestTryReceiveReportsReadyValue(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 7

	if got := tryReceive(ch); got != "received: 7" {
		t.Fatalf("tryReceive() = %q, want %q", got, "received: 7")
	}
}

func TestTryReceiveDoesNotBlockOnEmptyInput(t *testing.T) {
	result := make(chan string, 1)
	go func() {
		result <- tryReceive(make(chan int))
	}()

	select {
	case got := <-result:
		if got != "no value" {
			t.Fatalf("tryReceive() = %q, want %q", got, "no value")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("tryReceive() blocked on an empty channel")
	}
}
