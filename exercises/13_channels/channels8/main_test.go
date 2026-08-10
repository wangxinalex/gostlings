package main

import "testing"

func TestTryReceiveReportsReadyValue(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 7

	if got := tryReceive(ch); got != "received: 7" {
		t.Fatalf("tryReceive() = %q, want %q", got, "received: 7")
	}
}

func TestTryReceiveDoesNotBlockOnEmptyInput(t *testing.T) {
	if got := tryReceive(make(chan int)); got != "no value" {
		t.Fatalf("tryReceive() = %q, want %q", got, "no value")
	}
}
