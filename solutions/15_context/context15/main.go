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

type gatherOutcome struct {
	index int
	value int
	err   error
}

func gather(ctx context.Context, tasks []func(context.Context) (int, error)) ([]int, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	runCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	results := make([]int, len(tasks))
	outcomes := make(chan gatherOutcome, len(tasks))
	exited := make(chan struct{}, len(tasks))
	for index, task := range tasks {
		go func(index int, task func(context.Context) (int, error)) {
			defer func() { exited <- struct{}{} }()
			value, err := task(runCtx)
			outcomes <- gatherOutcome{index: index, value: value, err: err}
		}(index, task)
	}
	var firstErr error
	for range tasks {
		outcome := <-outcomes
		if outcome.err != nil && firstErr == nil {
			firstErr = outcome.err
			cancel()
		}
		results[outcome.index] = outcome.value
	}
	for range tasks {
		<-exited
	}
	if firstErr != nil {
		return nil, firstErr
	}
	return results, nil
}

func main() {}
