package main

import (
	"errors"
	"testing"
	"time"
)

func TestRetryWithTimerStopsAfterAFirstAttemptSuccess(t *testing.T) {
	previous := retryTimer
	timers := 0
	retryTimer = func(time.Duration) *time.Timer {
		timers++
		return time.NewTimer(0)
	}
	t.Cleanup(func() { retryTimer = previous })

	calls := 0
	attempts, err := retryWithTimer(4, time.Millisecond, func(int) error {
		calls++
		return nil
	})
	if err != nil {
		t.Fatalf("retryWithTimer() error = %v, want nil", err)
	}
	if attempts != 1 {
		t.Fatalf("retryWithTimer() attempts = %d, want 1", attempts)
	}
	if calls != 1 {
		t.Fatalf("work ran %d times, want 1", calls)
	}
	if timers != 0 {
		t.Fatalf("retryWithTimer() created %d timers, want none after a first-attempt success", timers)
	}
}

func TestRetryWithTimerStopsAtTheAttemptThatSucceeds(t *testing.T) {
	previous := retryTimer
	var delays []time.Duration
	retryTimer = func(delay time.Duration) *time.Timer {
		delays = append(delays, delay)
		return time.NewTimer(0)
	}
	t.Cleanup(func() { retryTimer = previous })

	const succeedOn = 3
	extra := make(chan int, 1)
	calls := 0
	attempts, err := retryWithTimer(5, 2*time.Millisecond, func(attempt int) error {
		calls++
		if attempt > succeedOn {
			select {
			case extra <- attempt:
			default:
			}
			return nil
		}
		if attempt < succeedOn {
			return errors.New("not yet")
		}
		return nil
	})
	if err != nil {
		t.Fatalf("retryWithTimer() error = %v, want nil", err)
	}
	if attempts != succeedOn {
		t.Fatalf("retryWithTimer() attempts = %d, want %d", attempts, succeedOn)
	}
	if calls != succeedOn {
		t.Fatalf("work ran %d times, want %d", calls, succeedOn)
	}
	if len(delays) != succeedOn-1 {
		t.Fatalf("retryWithTimer() waited %d times, want %d", len(delays), succeedOn-1)
	}
	for _, delay := range delays {
		if delay != 2*time.Millisecond {
			t.Fatalf("retryWithTimer() waited %v, want %v", delay, 2*time.Millisecond)
		}
	}
	select {
	case attempt := <-extra:
		t.Fatalf("work ran again as attempt %d after it had already succeeded", attempt)
	default:
	}
}

func TestRetryWithTimerReturnsTheLastErrorWhenAttemptsRunOut(t *testing.T) {
	previous := retryTimer
	timers := 0
	retryTimer = func(time.Duration) *time.Timer {
		timers++
		return time.NewTimer(0)
	}
	t.Cleanup(func() { retryTimer = previous })

	failures := []error{errors.New("first"), errors.New("second"), errors.New("third")}
	last := failures[len(failures)-1]
	attempts, err := retryWithTimer(len(failures), 2*time.Millisecond, func(attempt int) error {
		return failures[attempt-1]
	})
	if err != last {
		t.Fatalf("retryWithTimer() error = %v, want the last error %v", err, last)
	}
	if attempts != len(failures) {
		t.Fatalf("retryWithTimer() attempts = %d, want %d", attempts, len(failures))
	}
	if timers != len(failures)-1 {
		t.Fatalf("retryWithTimer() created %d timers, want %d", timers, len(failures)-1)
	}
}

func TestRetryWithTimerRejectsANonPositiveAttemptCount(t *testing.T) {
	previous := retryTimer
	timers := 0
	retryTimer = func(time.Duration) *time.Timer {
		timers++
		return time.NewTimer(0)
	}
	t.Cleanup(func() { retryTimer = previous })

	calls := 0
	attempts, err := retryWithTimer(0, time.Millisecond, func(int) error {
		calls++
		return nil
	})
	if err == nil {
		t.Fatal("retryWithTimer() with zero attempts returned a nil error")
	}
	if attempts != 0 {
		t.Fatalf("retryWithTimer() attempts = %d, want 0", attempts)
	}
	if calls != 0 || timers != 0 {
		t.Fatalf("retryWithTimer() ran work %d times and created %d timers, want none", calls, timers)
	}
}
