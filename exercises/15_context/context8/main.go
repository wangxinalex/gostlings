// Concept: context cancellation must unblock a blocked channel receive.
// Task: receive one value or return ctx.Err when the request is canceled.
// Hint: select between the input channel and ctx.Done(); return ctx.Err directly.
package main

import "context"

func receive(ctx context.Context, in <-chan int) (int, error) {
	// TODO: Make the blocked receive cancellation-aware.
	return 0, nil
}

func main() {}
