// Concept: an error in one worker must stop new work and still join every worker.
// Task: return the first observed job error, close stop once, and wait for all workers before returning.
// Expected behavior: a job with err stops the pool; a run without errors returns nil.
// Hint: separate the three responsibilities: jobs production, failure notification, and worker joining.
//       The producer sends jobs until the slice ends or stop closes, then closes its jobs channel.
//       A worker checks job.err before doing work; on the first error, send it to a capacity-one
//       buffered failure channel and exit. Successful workers continue until jobs closes or stop closes.
//       One coordinator receives the first failure, closes stop exactly once, and then waits for
//       every worker exit acknowledgement before returning the captured error. With no errors,
//       wait for the producer and all workers, then return nil.

package main

import "fmt"

type job struct {
	value int
	err   error
}

var onStopClosed = func() {}
var onWorkerExit = func() {}

func run(workers int, jobs []job) error {
	return nil // TODO: stop once on the first error and join all workers before returning it
}

func main() {
	fmt.Println(run(2, []job{{value: 1}, {value: 2}}))
}
