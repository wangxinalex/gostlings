// Concept: context.WithDeadline stops work at an absolute time.
// Task: derive a deadline context from ctx, defer cancel, and stop when it finishes.
// Hint: pass the caller's ctx to withDeadline; never replace it with Background.
package main

import (
	"context"
	"fmt"
	"time"
)

var withDeadline = context.WithDeadline
var workGate = make(chan struct{})
var runWork = func(ctx context.Context) string {
	select {
	case <-ctx.Done():
		return "work: deadline exceeded"
	case <-workGate:
		return "work: completed"
	}
}

func runUntil(ctx context.Context, deadline time.Time) string {
	// TODO: Use withDeadline(ctx, deadline), defer cancel, and select on Done.
	return "work: completed"
}

func main() {
	fmt.Println(runUntil(context.Background(), time.Now()))
}
