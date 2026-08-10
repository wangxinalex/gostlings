package main

import "testing"

func TestReceiveFastChoosesReadyInput(t *testing.T) {
	fast := make(chan string, 1)
	slow := make(chan string)
	fast <- "fast lane"

	if got := receiveFast(fast, slow); got != "fast lane" {
		t.Fatalf("receiveFast() = %q, want %q", got, "fast lane")
	}
}
