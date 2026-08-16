// Concept: a first error must stop new work, release every worker, and still have one owner for cancellation close.
// Task: process jobs with bounded workers and return the first job error after workers join.
// job carries value and err; a non-nil job.err is propagated and prevents later jobs from starting. Returned values
// are successful job results observed before cancellation.
// Hint: keep caller stop and internal cancel as separate signals. The producer selects on both before
//
//	sending each job and closes jobs when it stops. Workers select on both before receiving work,
//	before publishing a successful result, and while reporting an error.
//	A non-nil job.err goes to the capacity-one failure channel and causes that worker to exit.
//	The coordinator records the first observed error, closes internal cancel exactly once, drains
//	every worker exit acknowledgement, and only then returns. Successful results already observed
//	before cancellation may be kept; stop admitting later jobs once cancellation is observed.
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
