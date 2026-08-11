// Concept: context cancellation must unblock a blocked channel send.
// Task: send value or return ctx.Err when no receiver is ready and the request is canceled.
// Hint: select between out <- value and ctx.Done(); do not send without a cancellation case.
package main

import "context"

func send(ctx context.Context, out chan<- int, value int) error {
	// TODO: Make the blocked send cancellation-aware.
	return nil
}

func main() {}
