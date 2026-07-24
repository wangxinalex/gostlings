// Concept: table-driven tests with multiple cases
// Task: fill in the test cases so at least 3 pairs are tested, then run `go test` to verify
// Expected output: PASS (run with `go test ./exercises/14_testing/testing2 -v`)
// Hint: each case is a struct with inputs and the expected output; add cases for 0*5=0 and -2*3=-6 (Testing: go.dev/doc/tutorial/add-a-test)

package multiply

import "testing"

func TestMultiply(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{name: "2*3", a: 2, b: 3, want: 6},
		// TODO: Add at least two more test cases (e.g. 0*5, -2*3).
	}

	if len(cases) < 3 {
		t.Error("add at least 3 test cases to complete the table")
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Multiply(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("Multiply(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
