// Concept: cancellation propagates from a parent context to its children.
// Task: start count children derived from ctx and close the result only after all stop.
// Hint: derive every child from the parent context, release each derived context
//       on the way out, and wait for each child's Done channel before closing the
//       result.
// Stuck?: context.WithCancel returns the child plus its cancel function; defer it.

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
