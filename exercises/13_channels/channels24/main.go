// Problem: a worker pool can be blocked waiting for new jobs or blocked handing a
// result to a downstream consumer when cancellation arrives.
//
// Without this pattern: stopping the pool leaves workers alive and prevents the
// shared output from ever closing.
//
// Channels: jobs is receive-only input. stop requests cancellation. out carries
// results and is closed by the coordinator. exited confirms worker termination.
//
// Timeline:
//   select: jobs or stop
//   compute
//   select: out or stop
//   all workers exit -> close(out)
//
// Hint:
//   Use a stop case around the jobs receive and another around every result send.
//   Each worker reports one exit acknowledgement, including cancellation. The
//   coordinator waits for all acknowledgements before closing out.

package main

import "fmt"

var onSquareWorkerStart = func() {}
var onSquareWorkerBeforeSend = func() {}

func squareWorkers(stop <-chan struct{}, workers int, jobs <-chan int) <-chan int {
	return nil // TODO: use cancellable worker receives, sends, and coordinator-owned close(out)
}

func main() {
	jobs := make(chan int, 2)
	jobs <- 2
	jobs <- 3
	close(jobs)
	for result := range squareWorkers(make(chan struct{}), 2, jobs) {
		fmt.Println(result)
	}
}
