// Concept: a first error must stop new work, release every worker, and still have one owner for cancellation close.
// Task: process jobs with bounded workers and return the first job error after workers join.
// job carries value and err; a non-nil job.err is propagated and prevents later jobs from starting. Returned values
// are successful job results observed before cancellation.
// Hint: workers send errors to a capacity-one channel with a stop case. The coordinator alone closes its internal
// cancel channel once, stops the producer, drains worker exits, and only then returns the captured error.
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
