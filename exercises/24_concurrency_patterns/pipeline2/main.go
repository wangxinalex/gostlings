// Concept: context-aware pipeline stages and cancellation while sending
// Task: square every input, stop on cancellation, and close the output on every exit path
// Expected output: focused cancellation tests pass (run `go test ./exercises/24_concurrency_patterns/pipeline2`)
// Hint: select on ctx.Done() while receiving and again while sending; defer close(out)

package main

import "context"

func square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		// TODO: Receive from in until it closes. For each value, send its
		//       square, but let ctx.Done() interrupt a blocked send.
		for value := range in {
			out <- value * value
		}
	}()
	return out
}
