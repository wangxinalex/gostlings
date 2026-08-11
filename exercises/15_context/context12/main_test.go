package main

import (
	"context"
	"testing"
)

type lookupKey struct{}

func TestLookupPassesTheExactContext(t *testing.T) {
	ctx := context.WithValue(context.Background(), lookupKey{}, "request-42")
	got := lookup(ctx, func(received context.Context) string {
		if received != ctx {
			return "wrong context"
		}
		return received.Value(lookupKey{}).(string)
	})
	if got != "request-42" {
		t.Fatalf("lookup() = %q, want request-42", got)
	}
}

func TestLookupKeepsCancellationObservable(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	got := lookup(ctx, func(received context.Context) string {
		if received.Err() != context.Canceled {
			return "cancellation was hidden"
		}
		return "canceled"
	})
	if got != "canceled" {
		t.Fatalf("lookup() = %q, want canceled", got)
	}
}
