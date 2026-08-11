// Concept: graceful service shutdown stops intake before closing results.
// Task: process jobs until cancellation, close results, then close done.
// Hint: one goroutine owns both closures and checks ctx before each receive/send.
package main

import "context"

func serve(ctx context.Context, jobs <-chan int) (<-chan int, <-chan struct{}) {
	results := make(chan int)
	done := make(chan struct{})
	go func() {
		defer close(results)
		defer close(done)
		for {
			select {
			case <-ctx.Done():
				return
			case job, ok := <-jobs:
				if !ok {
					return
				}
				select {
				case results <- job * 2:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return results, done
}

func main() {}
