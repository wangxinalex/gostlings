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

func runPipeline(ctx context.Context, in <-chan job) (<-chan result, <-chan error) {
	// TODO: Process jobs, cancel on the first failure, and close both output channels.
	results := make(chan result)
	errorsOut := make(chan error, 1)
	return results, errorsOut
}

var errJobFailed = errors.New("job failed")
