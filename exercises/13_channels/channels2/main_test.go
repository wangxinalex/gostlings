package main

import (
	"io"
	"os"
	"testing"
	"time"
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

func TestOutput(t *testing.T) {
	gotCh := make(chan string, 1)
	go func() { gotCh <- captureStdout(main) }()
	select {
	case got := <-gotCh:
		const want = "1\n2\n"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("main() did not finish; the exercise is still broken (deadlock)")
	}
}
