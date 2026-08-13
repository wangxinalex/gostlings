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

func TestJoinWordsHasNoTrailingSpace(t *testing.T) {
	if got := joinWords([]string{"a", "b", "c"}); got != "a b c" {
		t.Fatalf("joinWords() = %q, want %q", got, "a b c")
	}
	if got := joinWords(nil); got != "" {
		t.Fatalf("joinWords(nil) = %q, want %q", got, "")
	}
}
