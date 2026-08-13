package main

import (
	"gostlings/internal/testutil"
	"testing"
)

func TestOutput(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	const want = "hello gostlings\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestNormalizeOnlyTrimsEdgesAndTabs(t *testing.T) {
	if got := normalize("\t  a  b \t"); got != "a  b" {
		t.Fatalf("normalize() = %q, want %q", got, "a  b")
	}
}
