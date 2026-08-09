package main

import (
	"gostlings/internal/testutil"
	"testing"
)

func TestOutput(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	const want = "42 is an int\nhello is a string\n3.14 is a float64\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
