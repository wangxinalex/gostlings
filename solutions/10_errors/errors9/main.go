// Concept: retrying only retryable errors
// Task: retry temporary failures, but stop immediately for non-retryable errors
// Expected output: request succeeded after retry
// Hint: attempts are numbered from 1. Call fetch once per attempt, return on nil,
//       and continue only when errors.Is(err, ErrTemporary) is true. If the loop
//       runs out of attempts, return a wrapped ErrTemporary to preserve its cause.

package main

import (
	"errors"
	"fmt"
)

var ErrTemporary = errors.New("temporary failure")

func fetch(attempt int) error {
	if attempt < 3 {
		return fmt.Errorf("attempt %d: %w", attempt, ErrTemporary)
	}
	return nil
}

func fetchWithRetry(maxAttempts int) error {
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		err := fetch(attempt)
		if err == nil {
			return nil
		}
		if !errors.Is(err, ErrTemporary) {
			return err
		}
	}
	return fmt.Errorf("request failed after %d attempts: %w", maxAttempts, ErrTemporary)
}

func main() {
	if err := fetchWithRetry(3); err == nil {
		fmt.Println("request succeeded after retry")
		return
	}
	fmt.Println("request failed")
}
