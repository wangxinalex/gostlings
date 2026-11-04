package main

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestRunServiceReturnsOrderedResponses(t *testing.T) {
	got, err := runService(context.Background(), 3, 2, []request{{value: 2}, {value: 1}, {value: 3}})
	if err != nil || len(got) != 3 || got[0].value != 4 || got[1].value != 2 || got[2].value != 6 {
		t.Fatalf("runService() = (%v, %v)", got, err)
	}
}

func TestRunServiceStopsOnFirstFailureAndCancellation(t *testing.T) {
	_, err := runService(context.Background(), 2, 1, []request{{value: 1}, {fail: true}, {value: 3}})
	if !errors.Is(err, errRequestFailed) {
		t.Fatalf("error = %v, want request failure", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	result := make(chan error, 1)
	go func() { _, err := runService(ctx, 2, 1, []request{{value: 1}}); result <- err }()
	select {
	case err := <-result:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("canceled error = %v", err)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runService() did not stop")
	}
}

func TestRunServiceBoundsActiveWorkAndJoinsWorkers(t *testing.T) {
	previous := serviceWork
	var active, peak, started, finished atomic.Int64
	release := make(chan struct{})
	finishGate := make(chan struct{})
	var releaseOnce, gateOnce sync.Once
	t.Cleanup(func() {
		releaseOnce.Do(func() { close(release) })
		gateOnce.Do(func() { close(finishGate) })
		serviceWork = previous
	})

	serviceWork = func(_ context.Context, req request) (int, error) {
		started.Add(1)
		current := active.Add(1)
		for {
			observed := peak.Load()
			if current <= observed || peak.CompareAndSwap(observed, current) {
				break
			}
		}
		<-release
		<-finishGate
		active.Add(-1)
		finished.Add(1)
		return req.value * 2, nil
	}

	requests := []request{{value: 1}, {value: 2}, {value: 3}, {value: 4}, {value: 5}}
	returned := make(chan struct{}, 1)
	go func() {
		runService(context.Background(), 5, 2, requests)
		returned <- struct{}{}
	}()

	deadline := time.After(500 * time.Millisecond)
	for started.Load() < 2 {
		select {
		case <-deadline:
			t.Fatalf("runService() started only %d requests, want 2", started.Load())
		default:
			time.Sleep(time.Millisecond)
		}
	}
	time.Sleep(10 * time.Millisecond)
	if got := started.Load(); got != 2 {
		t.Fatalf("%d requests active with limit 2, want exactly 2", got)
	}

	releaseOnce.Do(func() { close(release) })
	select {
	case <-returned:
		t.Fatal("runService() returned while workers were still blocked")
	case <-time.After(20 * time.Millisecond):
	}

	gateOnce.Do(func() { close(finishGate) })
	select {
	case <-returned:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runService() did not return after its workers finished")
	}
	if got := peak.Load(); got > 2 {
		t.Fatalf("peak concurrent work = %d, want at most 2", got)
	}
	if started.Load() != finished.Load() {
		t.Fatalf("runService() returned with %d workers still running", started.Load()-finished.Load())
	}
	if got := started.Load(); got != int64(len(requests)) {
		t.Fatalf("started %d requests, want %d", got, len(requests))
	}
}
