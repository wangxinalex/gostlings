// Concept: waiting for a dynamic number of jobs
// Task: run work once per job and return only after every job completes
// Expected behavior: every job contributes one result; an empty job slice returns an empty result slice.
// Hint: size the results slice from jobs, pass each index and job into its worker,
// and Wait after the loop. An empty slice needs no special wait.
// Version note: this repository uses sync.WaitGroup.Go, which requires Go 1.25+;
// older Go versions use Add, go, and defer Done. Do not rely on loop capture when
// writing code that should also work before Go 1.22.

package main

func runJobs(jobs []int, work func(int) string) []string {
	// TODO: Launch one worker for each job and wait for all results.
	return nil
}

func main() {}
