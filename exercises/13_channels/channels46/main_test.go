package main

import (
	"testing"
	"time"
)

func TestRateLimitStopsWhileWaitingForToken(t *testing.T) {
	stop := make(chan struct{})
	in := make(chan int, 1)
	in <- 4
	out := rateLimit(stop, make(chan struct{}), in)
	close(stop)
	closed46(t, out)
}

func TestRateLimitStopsBlockedOutputSendAfterConsumingToken(t *testing.T) {
	previous := onRateLimitBeforeSend
	beforeSend := make(chan struct{}, 1)
	onRateLimitBeforeSend = func() { beforeSend <- struct{}{} }
	t.Cleanup(func() { onRateLimitBeforeSend = previous })

	stop, in, tokens := make(chan struct{}), make(chan int), make(chan struct{})
	out := rateLimit(stop, tokens, in)
	inputSent := make(chan struct{})
	go func() { in <- 7; close(inputSent) }()
	select {
	case <-inputSent:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("rateLimit() did not receive input")
	}
	tokenSent := make(chan struct{})
	go func() { tokens <- struct{}{}; close(tokenSent) }()
	select {
	case <-tokenSent:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("rateLimit() did not consume its token")
	}
	select {
	case <-beforeSend:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("rateLimit() did not begin its blocked output send")
	}
	close(stop)
	closed46(t, out)
}

func closed46(t *testing.T, out <-chan int) {
	t.Helper()
	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("rateLimit() sent after stop")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("rateLimit() did not close after stop")
	}
}
