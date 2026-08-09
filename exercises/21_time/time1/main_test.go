package main

import (
	"gostlings/internal/testutil"
	"testing"
)

func TestOutput(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	const want = "Format: 2026-01-02 15:04:05\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
