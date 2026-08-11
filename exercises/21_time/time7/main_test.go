package main

import (
	"testing"
	"time"
)

func TestRunTickerStopsOnCancellationAndClosesDone(t *testing.T) {
	previous := tickerStopped
	stopped := make(chan struct{}, 1)
	tickerStopped = func() { stopped <- struct{}{} }
	t.Cleanup(func() { tickerStopped = previous })
	stop := make(chan struct{})
	ticks := make(chan time.Time)
	done := runTicker(stop, ticks)
	close(stop)
	select {
	case <-done:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runTicker() did not close done after stop")
	}
	select {
	case <-stopped:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runTicker() did not stop its ticker")
	}
}

func TestRunTickerStopsWhenTickInputCloses(t *testing.T) {
	previous := tickerStopped
	stopped := make(chan struct{}, 1)
	tickerStopped = func() { stopped <- struct{}{} }
	t.Cleanup(func() { tickerStopped = previous })
	ticks := make(chan time.Time)
	done := runTicker(make(chan struct{}), ticks)
	close(ticks)
	select {
	case <-done:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runTicker() did not close done after tick input closed")
	}
}
