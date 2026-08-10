// Concept: fuzz testing invariants and seed corpus
// Task: add a fuzz function that verifies reversing twice returns the original string
// Expected output: PASS (run with `go test ./exercises/14_testing/testing9`)
// Hint: keep the seed cases, then call f.Fuzz with a callback receiving *testing.T
//       and input string. Skip inputs for which utf8.ValidString is false because
//       this Unicode function intentionally defines behavior for valid UTF-8 text.
//       For valid input, compute Reverse(Reverse(input)) and compare it with input.
//       Try extra generated inputs with `go test -run=^$ -fuzz=FuzzReverse -fuzztime=2s`.

package reverse

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	f.Add("hello")
	f.Add("你好")
	f.Fuzz(func(t *testing.T, input string) {
		if !utf8.ValidString(input) {
			t.Skip()
		}
		if got := Reverse(Reverse(input)); got != input {
			t.Errorf("Reverse(Reverse(%q)) = %q, want %q", input, got, input)
		}
	})
}
