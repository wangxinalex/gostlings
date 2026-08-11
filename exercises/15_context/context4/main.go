// Concept: ctx.Err reports why a context stopped.
// Task: return ctx.Err after cancellation so callers can classify it.
// Hint: after ctx.Done() closes, return ctx.Err() directly; callers can use errors.Is.
package main

import "context"

func classify(ctx context.Context) error {
	// TODO: Return ctx.Err() once the context has finished.
	return nil
}

func main() {}
