package main

import (
	"gostlings/internal/testutil"
	"testing"
)

func TestOutput(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	const want = "odd\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
