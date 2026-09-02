// Problem: an application has many jobs but wants a fixed number of concurrent
// workers and one collected result slice.
//
// Without this pattern: one goroutine per job can consume too many resources;
// returning before workers finish loses results.
//
// Channels: jobsCh is owned by the producer and closes when the slice is sent.
// results carries values. exited carries one acknowledgement per worker. A
// coordinator owns close(results).
//
// Timeline:
//   producer sends jobs
//   producer closes jobsCh
//   workers finish and acknowledge exit
//   coordinator closes results
//   collector's range ends
//
// Hint:
//   Start the jobs producer in a goroutine and close jobsCh after the last send.
//   Start workers that range jobsCh and send to results. Collect results while
//   the pool runs; a coordinator waits for every worker acknowledgement and
//   closes results, allowing the collector's range to end.

package main

import "fmt"

var processJob = func(value int) int { return value * value }
var onWorkerExit = func() {}

func run(workers int, jobs []int) []int {
	return nil // TODO: produce jobs, fan out workers, and collect results after coordinator close
}

func main() {
	fmt.Println(run(2, []int{1, 2, 3, 4}))
}
