// Concept: an injected clock makes elapsed-time measurement deterministic and testable.
// Task: read the clock before and after the work and report how long the work took and
// whether it stayed within the fixed measureBudget.
// Expected behavior: ok reports whether the elapsed time is at most measureBudget, so the
// budget itself counts as within budget; the elapsed duration is exactly the difference
// between the two clock readings.
// Hint: call the injected now function once before the work and once after it, then compare
// the difference of the two readings against measureBudget.
// Stuck?: time.Time.Sub
package main

import "time"

const measureBudget = 100 * time.Millisecond

func measure(now func() time.Time, work func()) (time.Duration, bool) {
	// TODO: read the injected clock around work and compare the elapsed time with measureBudget.
	return 0, false
}

func main() {}
