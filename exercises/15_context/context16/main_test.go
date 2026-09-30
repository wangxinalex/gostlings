package main

import (
	"context"
	"errors"
	"testing"
)

type cleanupKey struct{}

func TestCleanupAfterRunsWhileTheCallerIsCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	ran := false
	err := cleanupAfter(ctx, func(cleanupCtx context.Context) error {
		if err := cleanupCtx.Err(); err != nil {
			t.Errorf("cleanup context error = %v, want a context detached from cancellation", err)
		}
		select {
		case <-cleanupCtx.Done():
			t.Error("cleanup context was canceled together with its parent")
		default:
		}
		ran = true
		return nil
	})
	if err != nil {
		t.Fatalf("cleanupAfter() error = %v, want nil", err)
	}
	if !ran {
		t.Fatal("cleanupAfter() did not run the cleanup work")
	}
}

func TestCleanupAfterKeepsValuesFromTheParentContext(t *testing.T) {
	parent := context.WithValue(context.Background(), cleanupKey{}, "kept")
	ctx, cancel := context.WithCancel(parent)
	cancel()

	var got any
	err := cleanupAfter(ctx, func(cleanupCtx context.Context) error {
		got = cleanupCtx.Value(cleanupKey{})
		return nil
	})
	if err != nil {
		t.Fatalf("cleanupAfter() error = %v, want nil", err)
	}
	if got != "kept" {
		t.Fatalf("cleanup context value = %v, want %q", got, "kept")
	}
}

func TestCleanupAfterReturnsTheWorkErrorUnchanged(t *testing.T) {
	failure := errors.New("cleanup failed")
	err := cleanupAfter(context.Background(), func(context.Context) error { return failure })
	if err != failure {
		t.Fatalf("cleanupAfter() error = %v, want the same error value %v", err, failure)
	}
}
