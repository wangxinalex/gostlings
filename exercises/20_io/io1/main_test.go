package main

import (
	"gostlings/internal/testutil"
	"io"
	"testing"
)

func TestOutput(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	const want = "Hello, io.Reader!\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
