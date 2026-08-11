// Concept: waiting for a dynamic number of jobs
// Task: run work once per job and return only after every job completes
// Expected behavior: every job contributes one result; an empty job slice returns an empty result slice.
// Hint: size the results slice from jobs, Add before each launch, and Wait after the loop. An empty slice needs no special wait.

package main

func runJobs(jobs []int, work func(int) string) []string {
	// TODO: Launch one worker for each job and wait for all results.
	return nil
}

func main() {}
