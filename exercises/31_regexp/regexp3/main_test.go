package main

import (
	"gostlings/internal/testutil"
	"testing"
)

func TestOutput(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	const want = "2026 08 13\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestParseDate(t *testing.T) {
	y, m, d := parseDate("2026-08-13")
	if y != "2026" || m != "08" || d != "13" {
		t.Fatalf("parseDate() = (%q, %q, %q), want (2026, 08, 13)", y, m, d)
	}
	if y, m, d = parseDate("bad"); y != "" || m != "" || d != "" {
		t.Fatalf("parseDate(\"bad\") = (%q, %q, %q), want empty", y, m, d)
	}
}
