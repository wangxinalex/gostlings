// Concept: cancellation propagates from a parent context to its children.
// Task: start count children derived from ctx and close the result only after all stop.
// Hint: each child should use context.WithCancel(ctx), defer its cancel function,
// and wait for child.Done().
package main

import "context"

var childStopped = func() {}
var childStarted = func() {}
var withCancel = context.WithCancel

func startChildren(ctx context.Context, count int) <-chan struct{} {
	// TODO: Start child contexts from ctx and close done after every child observes cancellation.
	done := make(chan struct{})
	close(done)
	return done
}

func main() {}
