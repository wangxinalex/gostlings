package testutil

import (
	"io"
	"os"
	"runtime"
	"testing"
	"time"
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

// TrackGoroutines records the current goroutine count and returns a function
// that fails the test when the program under test left goroutines running.
//
// Call it from the test goroutine and defer the returned function, so a leaked
// goroutine fails the test even when the program's own output looks right:
//
//	settle := testutil.TrackGoroutines(t)
//	defer settle()
//	got := testutil.CaptureStdout(t, main)
//
// Goroutines that are still winding down are given a short grace period, so the
// check reports leaks without depending on scheduler timing.
func TrackGoroutines(t *testing.T) func() {
	t.Helper()
	before := runtime.NumGoroutine()
	return func() {
		t.Helper()
		deadline := time.Now().Add(2 * time.Second)
		for {
			if after := runtime.NumGoroutine(); after <= before {
				return
			}
			if time.Now().After(deadline) {
				buffer := make([]byte, 1<<16)
				stack := buffer[:runtime.Stack(buffer, true)]
				t.Fatalf("the program left goroutines running: %d before, %d after\n%s",
					before, runtime.NumGoroutine(), stack)
			}
			time.Sleep(10 * time.Millisecond)
		}
	}
}
