package main

import (
	"io"
	"os"
	"testing"
)

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	os.Stdout = old
	b, _ := io.ReadAll(r)
	return string(b)
}

func TestOutputIncludesAllMapEntries(t *testing.T) {
	got := captureStdout(main)
	const want = "entries: 10000\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
