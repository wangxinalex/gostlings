// Concept: context.WithCancelCause preserves why cancellation happened.
// Task: cancel a derived context with requestCause and return that cause.
// Hint: context.Cause reads the sentinel supplied to CancelCauseFunc.
package main

import (
	"context"
	"errors"
)

var requestCause = errors.New("request rejected")
var withCancelCause = context.WithCancelCause

func runWithCause(ctx context.Context) error {
	// TODO: Derive a cause-aware context, cancel it with requestCause, and return context.Cause.
	return nil
}

func main() {}
