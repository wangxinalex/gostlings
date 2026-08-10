// Concept: isolated filesystem fixtures with t.TempDir
// Task: create a temporary directory, write the greeting there, and verify its contents
// Expected output: PASS (run with `go test ./exercises/14_testing/testing5`)
// Hint: t.TempDir returns a per-test directory that Go removes automatically.
//       Build a path with filepath.Join, call WriteGreeting, then use os.ReadFile
//       and compare the bytes with "hello\n". Do not write into the repository.

package greeting

import "testing"

func TestWriteGreeting(t *testing.T) {
	// TODO: Use t.TempDir and filepath.Join to create an isolated file path.
	//       Write the file, read it back, and assert that it contains "hello\n".
	t.Fatal("TODO: add a temporary-directory fixture")
}
