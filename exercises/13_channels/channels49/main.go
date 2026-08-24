// Problem: a bounded pool must distinguish caller cancellation from a job error,
// stop new work on the first error, and still release every worker.
// Without this pattern: one error can race with caller stop, multiple goroutines
// can close the same cancellation channel, or the function can return with work alive.
// Channels: caller stop and internal cancel are separate requests; jobs carries
// work; failure carries the first error; results carries successful values; exited
// joins workers. Only the coordinator closes internal cancel.
// Timeline: job error -> failure -> close(internal cancel) -> producer/workers exit -> join -> return
// Hint: keep caller stop and internal cancel separate. The producer selects on
// both before each job send and closes jobs when it stops. Workers select on both
// before receiving, publishing success, and reporting errors. Send the first
// non-nil job.err through a capacity-one failure channel; the coordinator records
// it, closes internal cancel once, drains every worker acknowledgement, and only
// then returns. Stop admitting later jobs after cancellation.
package main

import "fmt"

type job struct {
	value int
	err   error
}

var processFirstErrorBounded = func(value int) int { return value * value }

func runFirstErrorBounded(stop <-chan struct{}, workers int, jobs []job) ([]int, error) {
	return nil, nil // TODO: propagate the first error, cancel production once, and join every worker
}

func main() { fmt.Println(runFirstErrorBounded(make(chan struct{}), 1, nil)) }
