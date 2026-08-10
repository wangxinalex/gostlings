// Concept: deterministic environment-dependent tests with t.Setenv
// Task: set the environment only for this test and verify Greeting reads it
// Expected output: PASS (run with `go test ./exercises/14_testing/testing6`)
// Hint: use t.Setenv to set a test-only greeting; it restores the environment
//       automatically. Then assert Greeting returns the configured value.

package greeting

import "testing"

func TestGreetingFromEnvironment(t *testing.T) {
	// TODO: Set GOSTLINGS_GREETING for this test, then check Greeting().
	if got := Greeting(); got != "你好" {
		t.Errorf("Greeting() = %q, want %q", got, "你好")
	}
}
