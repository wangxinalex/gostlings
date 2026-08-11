package main

import (
	"context"
	"errors"
	"testing"
)

func TestRunWithCausePreservesTheSentinelCause(t *testing.T) {
	previous := withCancelCause
	var derived context.Context
	withCancelCause = func(parent context.Context) (context.Context, context.CancelCauseFunc) {
		derived, cancel := context.WithCancelCause(parent)
		return derived, cancel
	}
	t.Cleanup(func() { withCancelCause = previous })

	got := runWithCause(context.Background())
	if !errors.Is(got, requestCause) {
		t.Fatalf("runWithCause() = %v, want request cause", got)
	}
	if derived == nil || !errors.Is(context.Cause(derived), requestCause) {
		t.Fatalf("context.Cause(derived) = %v, want request cause", context.Cause(derived))
	}
}

func TestRunWithCauseRespectsParentCancellation(t *testing.T) {
	parent, cancel := context.WithCancel(context.Background())
	cancel()
	if err := runWithCause(parent); !errors.Is(err, context.Canceled) {
		t.Fatalf("runWithCause() = %v, want parent cancellation", err)
	}
}
