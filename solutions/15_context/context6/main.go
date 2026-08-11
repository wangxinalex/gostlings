// Concept: cancellation propagates from a parent context to its children.
// Task: start count children derived from ctx and close the result only after all stop.
// Hint: each child should use context.WithCancel(ctx), defer its cancel function,
// and wait for child.Done().
package main

import (
	"context"
	"sync"
)

var childStopped = func() {}
var withCancel = context.WithCancel

func startChildren(ctx context.Context, count int) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		var children sync.WaitGroup
		for range count {
			children.Add(1)
			go func() {
				defer children.Done()
				child, cancel := withCancel(ctx)
				defer cancel()
				<-child.Done()
				childStopped()
			}()
		}
		children.Wait()
	}()
	return done
}

func main() {}
