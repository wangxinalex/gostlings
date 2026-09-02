// Problem: one jobs stream needs concurrent processing, but callers still want
// one result stream and one clear completion point.
//
// Without this pattern: starting one goroutine per job has unbounded concurrency;
// letting workers close out independently can panic when several workers exit.
//
// Channels: jobs is caller-owned input. Its close means no more work. out is
// owned by this function and carries squares. exited carries one token per worker.
//
// Timeline:
//   jobs -> shared workers -> out -> caller
//   workers exit -> coordinator closes out
//
// Hint:
//   Start exactly workers goroutines in a loop. Each calls the start hook, ranges
//   jobs, sends each square to out, calls the exit hook, and reports one buffered
//   exited token. A coordinator receives every token and closes out once.
//   Return out immediately so the caller can range it while workers run.

package main

import "fmt"

var onSquareWorkerStart = func() {}
var onSquareWorkerExit = func() {}

func squareWorkers(workers int, jobs <-chan int) <-chan int {
	return nil // TODO: fan out jobs, fan in squares, and let one coordinator close out
}

func main() {
	jobs := make(chan int, 2)
	jobs <- 2
	jobs <- 3
	close(jobs)
	for result := range squareWorkers(2, jobs) {
		fmt.Println(result)
	}
}
