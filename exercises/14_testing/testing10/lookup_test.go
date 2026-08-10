// Concept: testing wrapped errors with errors.Is
// Task: assert the error category instead of comparing its formatted string
// Expected output: PASS (run with `go test ./exercises/14_testing/testing10`)
// Hint: call Find(42), first ensure err is non-nil, then use errors.Is(err, ErrNotFound).
//       The operation context in the message may change, but the category must remain stable.

package lookup

import (
	"errors"
	"testing"
)

func TestFindNotFound(t *testing.T) {
	_, err := Find(42)
	if err == nil {
		t.Fatal("Find() error = nil, want an error")
	}
	// TODO: Replace the newly-created sentinel below with ErrNotFound.
	if errors.Is(err, errors.New("record not found")) {
		return
	}
	t.Fatal("TODO: classify the wrapped error with errors.Is")
}
