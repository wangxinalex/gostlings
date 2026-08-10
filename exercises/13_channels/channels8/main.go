// Concept: worker pool — distribute jobs and close results after every worker exits
// Task: start workers, send all jobs, collect squared results, and shut down cleanly
// Expected behavior: one squared result per job; empty jobs also return promptly
// Hint: close jobs after sending, range over jobs in each worker, wait with sync.WaitGroup,
//       then close results from a coordinator

package main

func run(workers int, jobs []int) []int {
	// TODO: Create jobs/results channels, start workers, coordinate closure, and collect results.
	return nil
}
