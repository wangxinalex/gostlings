// Concept: retry policy must distinguish retryable errors and cancellable backoff.
// Task: retry only retryable errors and stop waiting for backoff after cancellation.
// Hint: select between each backoff event and ctx.Done; return permanent errors immediately.
package main

import (
	"context"
	"time"
)

type retryableError struct{}

func (retryableError) Error() string { return "retryable" }

func retry(ctx context.Context, attempts int, backoff <-chan time.Duration, work func() error) error {
	// TODO: Run work, retry only retryableError, and make backoff cancellation-aware.
	return nil
}
