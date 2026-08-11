package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestSendReturnsCancellationWhileOutputIsBlocked(t *testing.T) {
	out := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	result := make(chan error, 1)
	go func() { result <- send(ctx, out, 7) }()
	cancel()
	select {
	case err := <-result:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("send() error = %v, want context.Canceled", err)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("send() remained blocked after cancellation")
	}
}

func TestSendDeliversAValueWhenReceiverIsReady(t *testing.T) {
	out := make(chan int)
	result := make(chan error, 1)
	go func() { result <- send(context.Background(), out, 7) }()
	select {
	case value := <-out:
		if value != 7 {
			t.Fatalf("sent value = %d, want 7", value)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("send() did not deliver to a ready receiver")
	}
	if err := <-result; err != nil {
		t.Fatalf("send() error = %v, want nil", err)
	}
}
