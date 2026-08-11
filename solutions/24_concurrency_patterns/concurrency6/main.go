// Concept: a pipeline can carry typed results while a coordinator owns closure.
// Task: forward successful jobs, stop on the first failure, and report that error.
// Hint: use one result closer and a capacity-one error path; honor ctx on input/output.
package main

import (
	"context"
	"errors"
)

type job struct {
	value int
	fail  bool
}
type result struct {
	value int
	err   error
}

var errJobFailed = errors.New("job failed")

func runPipeline(ctx context.Context, in <-chan job) (<-chan result, <-chan error) {
	results := make(chan result)
	errorsOut := make(chan error, 1)
	go func() {
		defer close(results)
		defer close(errorsOut)
		for {
			select {
			case <-ctx.Done():
				return
			case item, ok := <-in:
				if !ok {
					return
				}
				if item.fail {
					errorsOut <- errJobFailed
					return
				}
				select {
				case results <- result{value: item.value * 2}:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return results, errorsOut
}

func main() {}
