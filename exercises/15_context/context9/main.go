// Concept: context cancellation must unblock a blocked channel send.
// Task: send value or return ctx.Err when no receiver is ready and the request is canceled.
// Hint: offer the value on one select case and the cancellation signal on the
//       other; never send without a cancellation case.
// Stuck?: select between the send and ctx.Done; return ctx.Err when canceled.

package main

import "context"

func send(ctx context.Context, out chan<- int, value int) error {
	// TODO: Make the blocked send cancellation-aware.
	return nil
}

func main() {}
