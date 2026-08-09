package testutil

import (
	"io"
	"os"
	"testing"
)

// CaptureStdout runs f and returns everything it writes to os.Stdout.
func CaptureStdout(t *testing.T, f func()) string {
	t.Helper()

	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w

	type readResult struct {
		data []byte
		err  error
	}
	readCh := make(chan readResult, 1)
	go func() {
		data, err := io.ReadAll(r)
		readCh <- readResult{data: data, err: err}
	}()

	defer func() {
		os.Stdout = old
		_ = w.Close()
		_ = r.Close()
	}()

	f()
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	os.Stdout = old

	result := <-readCh
	if result.err != nil {
		t.Fatal(result.err)
	}
	return string(result.data)
}
