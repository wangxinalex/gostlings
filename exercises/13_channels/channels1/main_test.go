package main

import (
	"gostlings/internal/testutil"
	"testing"
	"time"
)

func TestOutput(t *testing.T) {
	gotCh := make(chan string, 1)
	go func() { gotCh <- testutil.CaptureStdout(t, main) }()
	select {
	case got := <-gotCh:
		const want = "hi\n"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("main() did not finish; the exercise is still broken (deadlock)")
	}
}
