// Concept: worker pool — combine fan-out, result collection, and close ordering
// Task: process every job with a fixed number of workers and return all squares
// Expected behavior: one result per job; empty jobs return promptly
// Hint: close jobs after sending, wait for workers, then close results while the caller ranges results

package main

import "fmt"

func run(workers int, jobs []int) []int {
	// Thought: the producer closes jobs, which ends each worker's range. Only
	// after every worker exits may the coordinator close results, allowing the
	// caller's range to return the complete result set.
	return nil // TODO: build the jobs/results lifecycle
}

func main() {
	fmt.Println(run(2, []int{1, 2, 3, 4}))
}
