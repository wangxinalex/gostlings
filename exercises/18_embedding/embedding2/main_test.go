package main

import (
	"gostlings/internal/testutil"
	"testing"
)

func TestOutput(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	const want = "Alice (ID: 101)\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
