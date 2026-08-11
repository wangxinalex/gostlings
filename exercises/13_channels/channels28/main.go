// Concept: an error in one worker must stop new work and still join every worker.
// Task: return the first observed job error, close stop once, and wait for all workers before returning.
// Expected behavior: a job with err stops the pool; a run without errors returns nil.
// Hint: workers send errors to a buffered failure channel and send a separate exit acknowledgement when they finish.
//       One coordinator receives the first failure, closes stop once, then waits for every worker acknowledgement.

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
