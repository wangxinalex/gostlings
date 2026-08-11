// Concept: check cancellation before starting work.
// Task: run work exactly once only when ctx is still active, and report whether it started.
// Hint: use a non-blocking select on ctx.Done() before calling work.
package main

import "context"

func startIfActive(ctx context.Context, work func()) bool {
	// TODO: Reject an already-canceled context before invoking work.
	return false
}

func main() {}
