package main

import (
	"gostlings/internal/testutil"
	"os"
	"testing"
)

func TestOutput(t *testing.T) {
	t.Chdir(t.TempDir())
	t.Cleanup(func() { os.Remove("demo.txt") })
	got := testutil.CaptureStdout(t, main)
	const want = "hello, disk!\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
