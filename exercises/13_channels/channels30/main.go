// Problem: a fast producer can create more queued work than memory or the
// downstream system can handle.
// Without this pattern: an unbounded slice or one goroutine per job hides the
// overload instead of slowing the producer.
// Channels: jobsCh is the bounded work queue with exactly buffer slots; results
// carries squares; worker exits let a coordinator close results.
// Timeline: producer fills queue -> queue full blocks producer -> worker frees slot -> producer continues
// Hint: create jobsCh with `make(chan int, buffer)`. Send every job from a
// producer goroutine and close jobsCh afterward. Collect results concurrently;
// the queue capacity, not len(jobs), controls buffered work.

package main

import "fmt"

var processBoundedJob = func(value int) int { return value * value }
var onBoundedQueue = func(capacity int) {}
var onBoundedProcessStart = func(value int) {}

func runBounded(workers, buffer int, jobs []int) []int {
	return nil // TODO: use a bounded jobs channel, worker acknowledgements, and coordinator-owned result close
}

func main() {
	fmt.Println(runBounded(2, 1, []int{1, 2, 3, 4}))
}
