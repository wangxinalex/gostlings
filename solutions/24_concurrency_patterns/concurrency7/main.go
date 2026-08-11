// Concept: batching combines a size threshold, explicit flush events, and cancellation.
// Task: emit full batches immediately and partial batches on flush; discard on cancel.
// Hint: keep one owner for the current slice and close output on every exit.
package main

import (
	"context"
	"time"
)

func batch(ctx context.Context, in <-chan int, flush <-chan time.Time, size int) <-chan []int {
	out := make(chan []int)
	go func() {
		defer close(out)
		current := make([]int, 0, size)
		emit := func() bool {
			if len(current) == 0 {
				return true
			}
			copyOf := append([]int(nil), current...)
			select {
			case out <- copyOf:
				current = current[:0]
				return true
			case <-ctx.Done():
				return false
			}
		}
		for {
			select {
			case <-ctx.Done():
				return
			case value, ok := <-in:
				if !ok {
					return
				}
				current = append(current, value)
				if len(current) >= size && !emit() {
					return
				}
			case <-flush:
				if !emit() {
					return
				}
			}
		}
	}()
	return out
}

func main() {}
