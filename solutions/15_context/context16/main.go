// Concept: cleanup work must survive the cancellation that triggered it.
// Task: run work on a context detached from the caller's cancellation but still carrying
// the caller's values, and report the work's error unchanged.
// Expected behavior: cleanup still runs when ctx is already canceled; values attached to
// ctx stay visible to the cleanup context; the cleanup error is returned unchanged.
// Hint: derive a detaching context from the caller's context, so the values survive while
// its cancellation does not, then run the work on it and return the work's error directly.
// Stuck?: context.WithoutCancel
package main

import "context"

func cleanupAfter(ctx context.Context, work func(context.Context) error) error {
	cleanupCtx := context.WithoutCancel(ctx)
	return work(cleanupCtx)
}

func main() {}
