// Concept: a pipeline stage must observe cancellation while receiving and sending.
// Task: square values and close output on normal completion or cancellation.
// Hint: select on ctx.Done for both channel operations and defer close(out).
package main

import "context"

func square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case value, ok := <-in:
				if !ok {
					return
				}
				select {
				case <-ctx.Done():
					return
				case out <- value * value:
				}
			}
		}
	}()
	return out
}

func main() {}
