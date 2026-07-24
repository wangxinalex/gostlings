// Concept: subtests with t.Run for grouped test output
// Task: write two subtests using t.Run — one for "even" and one for "odd" — so `go test -v` shows named subtests
// Expected output: both subtests PASS (run with `go test ./exercises/14_testing/testing3 -v`)
// Hint: t.Run("even", func(t *testing.T) { ... }) creates a named subtest (Testing: go.dev/doc/tutorial/add-a-test)

package even

import "testing"

func TestIsEven(t *testing.T) {
	// TODO: Write two t.Run subtests:
	//       "even" — IsEven(4) should be true
	//       "odd"  — IsEven(3) should be false
	t.Error("write two t.Run subtests as described in the TODO")
}
