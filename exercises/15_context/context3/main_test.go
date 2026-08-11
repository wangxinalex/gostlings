package main

import (
	"context"
	"testing"
)

func TestHandlerReturnsTypedUserValue(t *testing.T) {
	ctx := context.WithValue(context.Background(), userKey{}, "Alice")
	if got := handler(ctx); got != "user: Alice" {
		t.Fatalf("handler() = %q, want %q", got, "user: Alice")
	}
}

func TestHandlerReturnsFallbackForMissingOrWrongValue(t *testing.T) {
	if got := handler(context.Background()); got != "user: guest" {
		t.Fatalf("handler() with no value = %q, want %q", got, "user: guest")
	}
	wrongType := context.WithValue(context.Background(), userKey{}, 42)
	if got := handler(wrongType); got != "user: guest" {
		t.Fatalf("handler() with wrong value type = %q, want %q", got, "user: guest")
	}
}
