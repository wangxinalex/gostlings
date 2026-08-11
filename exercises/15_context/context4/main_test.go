package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestClassifyCanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := classify(ctx); !errors.Is(err, context.Canceled) {
		t.Fatalf("classify() error = %v, want context.Canceled", err)
	}
}

func TestClassifyDeadlineExceededContext(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(-time.Second))
	defer cancel()
	if err := classify(ctx); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("classify() error = %v, want context.DeadlineExceeded", err)
	}
}
