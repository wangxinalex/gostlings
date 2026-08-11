package main

import (
	"testing"
	"time"
)

func TestReceiveFastChoosesReadyInput(t *testing.T) {
	t.Run("fast is ready", func(t *testing.T) {
		fast := make(chan string, 1)
		slow := make(chan string)
		fast <- "fast lane"

		if got := receiveFast(fast, slow); got != "fast lane" {
			t.Fatalf("receiveFast() = %q, want %q", got, "fast lane")
		}
	})

	t.Run("slow is ready", func(t *testing.T) {
		fast := make(chan string)
		slow := make(chan string, 1)
		slow <- "slow lane"

		result := make(chan string, 1)
		go func() {
			result <- receiveFast(fast, slow)
		}()

		select {
		case got := <-result:
			if got != "slow lane" {
				t.Fatalf("receiveFast() = %q, want %q", got, "slow lane")
			}
		case <-time.After(200 * time.Millisecond):
			// Unblock an implementation that is incorrectly waiting only on fast,
			// so the failing test does not leave a goroutine behind.
			close(fast)
			t.Fatal("receiveFast blocked even though slow was ready")
		}
	})
}
