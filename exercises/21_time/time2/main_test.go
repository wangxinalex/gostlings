package main

import (
	"gostlings/internal/testutil"
	"testing"
)

func TestOutput(t *testing.T) {
	settle := testutil.TrackGoroutines(t)
	defer settle()
	got := testutil.CaptureStdout(t, main)
	const want = "tick\ntick\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
