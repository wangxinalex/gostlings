package main

import (
	"context"
	"errors"
	"testing"
)

func TestSubmitEnqueuesWhenCapacityIsAvailable(t *testing.T) {
	queue := make(chan submitJob, 1)
	if err := submit(context.Background(), queue, make(chan submitResult), 7); err != nil {
		t.Fatalf("submit() error = %v", err)
	}
	select {
	case job := <-queue:
		if job.value != 7 {
			t.Fatalf("job value = %d, want 7", job.value)
		}
	default:
		t.Fatal("submit() did not enqueue a job")
	}
}

func TestSubmitRejectsFullQueueAndCancellation(t *testing.T) {
	queue := make(chan submitJob, 1)
	queue <- submitJob{value: 1}
	if err := submit(context.Background(), queue, nil, 2); !errors.Is(err, errQueueFull) {
		t.Fatalf("full queue error = %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := submit(ctx, make(chan submitJob), nil, 3); !errors.Is(err, context.Canceled) {
		t.Fatalf("canceled error = %v", err)
	}
}
