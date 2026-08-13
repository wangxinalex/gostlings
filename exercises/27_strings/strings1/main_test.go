package main

import (
	"gostlings/internal/testutil"
	"testing"
)

func TestOutput(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	const want = "true\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestContainsReportsAbsenceAndPresence(t *testing.T) {
	if contains("gostlings", "rust") {
		t.Fatal("contains() reported a missing substring")
	}
	if !contains("gostlings", "ost") {
		t.Fatal("contains() missed a present substring")
	}
}
