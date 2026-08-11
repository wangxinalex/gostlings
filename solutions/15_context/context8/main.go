// Concept: context cancellation must unblock a blocked channel receive.
// Task: receive one value or return ctx.Err when the request is canceled.
// Hint: select between the input channel and ctx.Done(); return ctx.Err directly.
package main

import "context"

func receive(ctx context.Context, in <-chan int) (int, error) {
	select {
	case value := <-in:
		return value, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

func main() {}
