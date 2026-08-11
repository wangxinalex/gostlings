// Concept: atomic metrics are read after workers join.
// Task: count completed and canceled jobs without a data race.
// Hint: use atomic counters and make each worker observe ctx before reporting its outcome.
package main

import "context"

var measureWork = func(ctx context.Context, job int) bool { return ctx.Err() == nil }

func runMeasured(ctx context.Context, workers int, jobs []int) (completed, canceled int64) {
	// TODO: Run jobs, atomically count outcomes, and join all workers before returning.
	return 0, 0
}
