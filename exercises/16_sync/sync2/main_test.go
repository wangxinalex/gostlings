package main

import (
	"gostlings/internal/testutil"
	"testing"
)

func TestOutputIncludesAllMapEntries(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	const want = "entries: 10000\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
