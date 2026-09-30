// Concept: context cancellation must unblock a blocked channel receive.
// Task: receive one value or return ctx.Err when the request is canceled.
// Hint: wait for the input and the cancellation signal at the same time, and
//       return the context's error when cancellation wins.
// Stuck?: select between the input channel and ctx.Done; return ctx.Err.

package main

import "context"

func receive(ctx context.Context, in <-chan int) (int, error) {
	// TODO: Make the blocked receive cancellation-aware.
	return 0, nil
}

func main() {}
