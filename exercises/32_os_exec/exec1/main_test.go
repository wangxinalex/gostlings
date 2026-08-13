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

func TestEcho(t *testing.T) {
	got, err := echo("hello", "gostlings")
	if err != nil {
		t.Fatal(err)
	}
	if got != "hello gostlings" {
		t.Fatalf("echo() = %q, want %q", got, "hello gostlings")
	}
}
