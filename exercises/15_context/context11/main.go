// Concept: check cancellation before starting work.
// Task: run work exactly once only when ctx is still active, and report whether it started.
// Hint: test the cancellation signal without blocking before running the work;
//       a canceled context means the work must not start.
// Stuck?: a select on ctx.Done with a default branch.

package main

import "context"

func startIfActive(ctx context.Context, work func()) bool {
	// TODO: Reject an already-canceled context before invoking work.
	return false
}

func main() {}
