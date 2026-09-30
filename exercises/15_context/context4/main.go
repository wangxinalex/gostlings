// Concept: ctx.Err reports why a context stopped.
// Task: return ctx.Err after cancellation so callers can classify it.
// Hint: once the context's Done channel is closed, hand back the context's own
//       error value so callers can classify it with errors.Is.
// Stuck?: ctx.Err reports Canceled or DeadlineExceeded.

package main

import "context"

func classify(ctx context.Context) error {
	// TODO: Return ctx.Err() once the context has finished.
	return nil
}

func main() {}
