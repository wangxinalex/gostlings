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
	derived, cancel := withCancelCause(ctx)
	defer cancel(nil)
	cancel(requestCause)
	return context.Cause(derived)
}

func main() {}
