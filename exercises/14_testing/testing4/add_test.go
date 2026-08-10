// Concept: test helpers and useful failure locations
// Task: turn assertEqual into a reusable test helper and use it for every table case
// Expected output: PASS (run with `go test ./exercises/14_testing/testing4`)
// Hint: call t.Helper() as the first statement so failures point to the test case,
//       not the helper. Then compare got and want and report a useful t.Errorf.

package add

import "testing"

func assertEqual(t *testing.T, got, want int) {
	// TODO: Mark this function as a helper and report a mismatch.
	t.Error("TODO: implement assertEqual")
}

func TestAdd(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{name: "positive", a: 2, b: 3, want: 5},
		{name: "zero", a: 0, b: 4, want: 4},
		{name: "negative", a: -2, b: 3, want: 1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assertEqual(t, Add(tc.a, tc.b), tc.want)
		})
	}
}
