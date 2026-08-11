// Concept: worker pools need cancellation around both jobs receives and result sends.
// Task: square jobs with workers workers until jobs closes or stop closes.
// Expected behavior: stop releases workers blocked waiting for a job or waiting for a result receiver.
// Hint: select on stop when receiving jobs, and select on stop again when sending each square to out.
//       Workers send exit acknowledgements; a coordinator waits for all of them before close(out).

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
