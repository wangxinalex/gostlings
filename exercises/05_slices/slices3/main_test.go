package main

import (
	"gostlings/internal/testutil"
	"testing"
)

func TestOutput(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	const want = "[2 3 4]\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
