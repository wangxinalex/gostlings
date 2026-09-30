package main

import (
	"gostlings/internal/testutil"
	"testing"
)

func TestOutput(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	const want = "go is fun\nclosed: true\n3 2 1\nclosed: true\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// tokens satisfies ReadCloser from outside main.go: any type with both methods works.
type tokens struct {
	values []string
	index  int
	closed bool
}

func (t *tokens) Read() (string, bool) {
	if t.index >= len(t.values) {
		return "", false
	}
	value := t.values[t.index]
	t.index++
	return value, true
}

func (t *tokens) Close() error {
	t.closed = true
	return nil
}

func TestDrainAndCloseAcceptsAnyReadCloser(t *testing.T) {
	stream := &tokens{values: []string{"alpha", "beta"}}
	if got := drainAndClose(stream); got != "alpha beta" {
		t.Fatalf("drainAndClose() = %q, want %q", got, "alpha beta")
	}
	if !stream.closed {
		t.Fatal("drainAndClose did not close the reader")
	}
}
