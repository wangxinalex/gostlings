// Concept: fakes and dependency injection
// Task: configure a fake UserStore and test Welcome without a database or network
// Expected output: PASS (run with `go test ./exercises/14_testing/testing7`)
// Hint: the fake only needs to implement Name. Give it a name field, create it
//       with name "Ada", call Welcome with context.Background(), and assert the
//       returned string is "welcome, Ada" and the error is nil.

package welcome

import (
	"context"
	"testing"
)

type fakeStore struct {
	name string
}

func (f fakeStore) Name(context.Context, int) (string, error) {
	return f.name, nil
}

func TestWelcome(t *testing.T) {
	store := fakeStore{}
	got, err := Welcome(context.Background(), store, 42)
	if err != nil {
		t.Fatalf("Welcome() error = %v", err)
	}
	// TODO: Configure the fake so this assertion receives "welcome, Ada".
	if got != "welcome, Ada" {
		t.Errorf("Welcome() = %q, want %q", got, "welcome, Ada")
	}
}
