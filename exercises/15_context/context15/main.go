// Concept: concurrent fan-out must cancel the rest and join every goroutine on the first error.
// Task: run every task concurrently, return the values in task order, cancel the remaining
// work when one task fails, and wait for every started goroutine before returning.
// Expected behavior: all tasks succeed -> ordered values and a nil error; one task fails ->
// that error, returned only after every goroutine exited; an already canceled ctx -> the
// context error and no task started; no tasks -> an empty non-nil slice and a nil error.
// Hint: derive a cancelable context, give every goroutine one slot in a buffered result
// channel plus a completion signal, cancel as soon as the first error arrives, and count
// completions before returning.
// Stuck?: context.WithCancel
package main

import "context"

func gather(ctx context.Context, tasks []func(context.Context) (int, error)) ([]int, error) {
	// TODO: fan out over tasks, keep task order, cancel the rest on the first error, and join all.
	return nil, nil
}

func main() {}
