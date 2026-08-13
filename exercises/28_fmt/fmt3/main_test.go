package main

import (
	"bytes"
	"gostlings/internal/testutil"
	"testing"
)

func TestOutput(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	const want = "Hello, Ada!\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGreetWritesToWriter(t *testing.T) {
	var buf bytes.Buffer
	if err := greet(&buf, "Ada"); err != nil {
		t.Fatal(err)
	}
	if got := buf.String(); got != "Hello, Ada!" {
		t.Fatalf("greet() = %q, want %q", got, "Hello, Ada!")
	}
}
