// Problem: one failed job should stop admitting new work, but other workers may
// still be receiving or processing jobs.
// Without this pattern: continuing after failure wastes work; returning
// immediately can leak workers and race with later sends.
// Channels: jobs carries work; failure is a capacity-one first-error signal;
// stop broadcasts cancellation; exited joins workers. The coordinator owns the
// one close(stop) and returns only after every worker exits.
// Timeline: worker observes error -> failure -> coordinator closes stop -> producer/workers exit -> join -> return error
// Hint: keep production, failure reporting, and joining separate. The producer
// stops sending when stop closes and closes jobs. A worker reports the first
// error and exits; successful workers select on jobs or stop. The coordinator
// captures the first failure, closes stop once, drains every exit acknowledgement,
// and returns the error only after cleanup.

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
