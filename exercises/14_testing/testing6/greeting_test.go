// Concept: deterministic environment-dependent tests with t.Setenv
// Task: set the environment only for this test and verify Greeting reads it
// Expected output: PASS (run with `go test ./exercises/14_testing/testing6`)
// Hint: t.Setenv("GOSTLINGS_GREETING", "你好") changes the variable for this
//       test and restores it automatically. Then assert Greeting() returns "你好".

package greeting

import "testing"

func TestGreetingFromEnvironment(t *testing.T) {
	// TODO: Set GOSTLINGS_GREETING for this test, then check Greeting().
	if got := Greeting(); got != "你好" {
		t.Errorf("Greeting() = %q, want %q", got, "你好")
	}
}
