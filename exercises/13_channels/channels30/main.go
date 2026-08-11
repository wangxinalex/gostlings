// Concept: a bounded jobs channel provides backpressure without allocating one queue slot per job.
// Task: square jobs with workers workers while the internal jobs queue has exactly buffer slots.
// Expected behavior: workers and submission make progress; no extra queue grows with len(jobs).
// Hint: make jobsCh with make(chan int, buffer), then have a producer send every job and close jobsCh.
//       Start result collection while the producer and workers run; a coordinator closes results after workers exit.

package main

import "fmt"

var processBoundedJob = func(value int) int { return value * value }
var onBoundedQueue = func(capacity int) {}

func runBounded(workers, buffer int, jobs []int) []int {
	return nil // TODO: use a bounded jobs channel, worker acknowledgements, and coordinator-owned result close
}

func main() {
	fmt.Println(runBounded(2, 1, []int{1, 2, 3, 4}))
}
