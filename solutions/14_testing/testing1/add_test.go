// Concept: writing a basic table-driven test
// Task: the expected value in this test is wrong; fix it so `go test` passes
// Expected output: PASS (run with `go test ./exercises/14_testing/testing1`)
// Hint: tests are run with `go test`; the expected value for 2 + 3 is 5 (Testing: go.dev/doc/tutorial/add-a-test)

package add

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5
	if got != want {
		t.Errorf("Add(2, 3) = %d; want %d", got, want)
	}
}
