package main

import (
	"gostlings/internal/testutil"
	"testing"
)

func TestOutput(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	const want = "[Ada       ] $3.50\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestPadNameAndFormatPrice(t *testing.T) {
	if got := padName("Ada"); got != "Ada       " {
		t.Fatalf("padName() = %q, want %q", got, "Ada       ")
	}
	if got := formatPrice(3.5); got != "3.50" {
		t.Fatalf("formatPrice() = %q, want %q", got, "3.50")
	}
}
