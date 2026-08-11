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
	done := make(chan struct{})
	exited := make(chan struct{})

	for range count {
		go func() {
			child, cancel := withCancel(ctx)
			defer func() {
				cancel()
				exited <- struct{}{}
			}()
			childStarted()
			<-child.Done()
			childStopped()
		}()
	}

	go func() {
		defer close(done)
		for range count {
			<-exited
		}
	}()
	return done
}

func main() {}
