// Concept: cancellation must travel through every stage of a pipeline.
// Task: multiply each value by two, add one in a second stage, and close output on cancel.
// Hint: each stage needs cancellation-aware receive and send cases.
package main

import "context"

func pipeline(ctx context.Context, in <-chan int) <-chan int {
	doubled := make(chan int)
	go func() {
		defer close(doubled)
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
				case doubled <- value * 2:
				}
			}
		}
	}()
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case value, ok := <-doubled:
				if !ok {
					return
				}
				select {
				case <-ctx.Done():
					return
				case out <- value + 1:
				}
			}
		}
	}()
	return out
}

func main() {}
