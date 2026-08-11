// Concept: completing each started goroutine exactly once
// Task: visit every job and return the number of workers that completed
// Expected behavior: every job is visited once and the returned count equals the number of jobs.
// Hint: use one result slot per job. Set it after visit returns, defer Done in the worker, then count after Wait.

package main

func runEach(jobs []int, visit func(int)) int {
	// TODO: Start one worker per job, join them, and count completed visits.
	return 0
}

func main() {}
