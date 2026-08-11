package main

import (
	"testing"
	"time"
)

func TestAwaitOrCancelReturnsResult(t *testing.T) {
	result := make(chan string, 1)
	result <- "ready"
	if got := awaitOrCancel(result, time.Second); got != "ready" {
		t.Fatalf("awaitOrCancel() = %q, want ready", got)
	}
}

func TestAwaitOrCancelReportsClosedResult(t *testing.T) {
	result := make(chan string)
	close(result)
	if got := awaitOrCancel(result, 1000000000); got != "canceled" {
		t.Fatalf("awaitOrCancel() = %q, want canceled", got)
	}
}

func TestAwaitOrCancelReportsTimeout(t *testing.T) {
	if got := awaitOrCancel(make(chan string), 1); got != "timed out" {
		t.Fatalf("awaitOrCancel() = %q, want timed out", got)
	}
}
