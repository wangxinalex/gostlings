// Concept: isolated filesystem fixtures with t.TempDir
// Task: create a temporary directory, write the greeting there, and verify its contents
// Expected output: PASS (run with `go test ./exercises/14_testing/testing5`)
// Hint: t.TempDir returns a per-test directory that Go removes automatically.
//       Build a path with filepath.Join, call WriteGreeting, then use os.ReadFile
//       and compare the bytes with "hello\n". Do not write into the repository.

package greeting

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteGreeting(t *testing.T) {
	path := filepath.Join(t.TempDir(), "greeting.txt")
	if err := WriteGreeting(path); err != nil {
		t.Fatalf("WriteGreeting() error = %v", err)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}
	if string(got) != "hello\n" {
		t.Errorf("file contains %q, want %q", got, "hello\n")
	}
}
