// Concept: completing each started goroutine exactly once
// Task: visit every job and return the number of workers that completed
// Expected behavior: every job is visited once and the returned count equals the number of jobs.
// Hint: use one result slot per job. Pass both the result index and job into the
// worker, set the slot after visit returns, defer Done, then count after Wait.
// Explicitly passing the index avoids relying on Go 1.22's new range-variable scope.

package main

func runEach(jobs []int, visit func(int)) int {
	// TODO: Start one worker per job, join them, and count completed visits.
	return 0
}

func main() {}
