package main

import (
	"gostlings/internal/testutil"
	"testing"
)

func TestOutput(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	const want = "[[1 2 3] [4 5 6] [7]]\n[[1 2] [3 4] [5 6] [7]]\nsize 0 is nil: true\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
