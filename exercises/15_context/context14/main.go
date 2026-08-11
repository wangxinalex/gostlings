// Concept: a request combines deadline propagation, cooperative workers, and a final join.
// Task: run workers under a timeout, return one summary on success, and return the
// context error only after every worker has stopped.
// Hint: derive one timeout context, select on ctx.Done in every worker, and join
// workers through an acknowledgment channel before returning.
package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var withRequestTimeout = context.WithTimeout
var requestWork = func(ctx context.Context, id int) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
		return fmt.Sprintf("request-%d", id), nil
	}
}

func runRequest(ctx context.Context, workers int) (string, error) {
	// TODO: Derive a timeout context, run workers, and join every worker before returning.
	if workers < 1 {
		return "", errors.New("workers must be positive")
	}
	_ = time.Second
	return "", nil
}

func main() {}
