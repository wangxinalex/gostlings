// Concept: buffering changes decoupling, not cancellation ownership.
// Task: double values through a buffered stage and close output on cancellation.
// Hint: make the buffer size explicit and retain cancellation cases in both directions.
package main

import "context"

func bufferedPipeline(ctx context.Context, in <-chan int, buffer int) <-chan int {
	if buffer < 0 {
		buffer = 0
	}
	out := make(chan int, buffer)
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
				case out <- value * 2:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return out
}

func main() {}
