// Concept: retry policy must distinguish retryable errors and cancellable backoff.
// Task: retry only retryable errors and stop waiting for backoff after cancellation.
// Hint: select between each backoff event and ctx.Done; return permanent errors immediately.
package main

import (
	"context"
	"errors"
	"time"
)

type retryableError struct{}

func (retryableError) Error() string { return "retryable" }

func retry(ctx context.Context, attempts int, backoff <-chan time.Duration, work func() error) error {
	if attempts < 1 {
		return errors.New("attempts must be positive")
	}
	for attempt := 0; attempt < attempts; attempt++ {
		if err := work(); err != nil {
			var retryable retryableError
			if !errors.As(err, &retryable) {
				return err
			}
			if attempt == attempts-1 {
				return err
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case delay, ok := <-backoff:
				if !ok {
					return err
				}
				if delay > 0 {
					timer := time.NewTimer(delay)
					select {
					case <-ctx.Done():
						timer.Stop()
						return ctx.Err()
					case <-timer.C:
					}
				}
			}
			continue
		}
		return nil
	}
	return nil
}

func main() {}
