package testutil

import (
	"io"
	"os"
	"testing"
)

// CaptureStdout runs f and returns everything it writes to os.Stdout.
func CaptureStdout(t *testing.T, f func()) string {
	t.Helper()

	// Save the original stdout so the test cannot affect later tests.
	old := os.Stdout
	// Use a pipe to collect everything written by the function under test.
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
	// Read concurrently so a large amount of output cannot fill the pipe buffer
	// and block the function under test.
	go func() {
		data, err := io.ReadAll(r)
		readCh <- readResult{data: data, err: err}
	}()

	defer func() {
		// Restore stdout and close both ends of the pipe on every exit path.
		os.Stdout = old
		_ = w.Close()
		_ = r.Close()
	}()

	f()
	// Closing the writer signals EOF to the reader goroutine.
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	// Restore stdout before waiting for the reader to finish.
	os.Stdout = old

	result := <-readCh
	if result.err != nil {
		t.Fatal(result.err)
	}
	return string(result.data)
}
