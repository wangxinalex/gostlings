// Problem: an API should prevent callers from performing channel operations that
// violate ownership.
// Without this pattern: returning bidirectional channels lets callers receive
// from a jobs queue or close a result stream owned by the pool.
// Channels: jobs is returned as chan<- int so callers can send and close it;
// results is returned as <-chan int so callers can only receive and range it.
// The pool owns worker exits and close(results).
// Timeline: caller sends jobs -> caller closes jobs -> workers exit -> coordinator closes results -> caller ranges results
// Hint: create both channels internally, return jobs as `chan<- int` and results
// as `<-chan int`. Workers range jobs; one coordinator waits for all exit tokens
// before closing results exactly once.

package main

import "fmt"

var onPoolWorkerStart = func() {}
var onPoolWorkerExit = func() {}

func startPool(workers int) (chan<- int, <-chan int) {
	return nil, nil // TODO: expose directional handles and coordinate result closure
}

func main() {
	jobs, results := startPool(2)
	jobs <- 2
	jobs <- 3
	close(jobs)
	for result := range results {
		fmt.Println(result)
	}
}
