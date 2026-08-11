// Concept: add all known work before waiting for a worker group.
// Task: apply runTask to every job concurrently and return the sum of its results.
// Hint: use a buffered result channel and a WaitGroup; never call Wait before all Add calls.
package main

var runTask = func(job int) int { return job }

func runTasks(jobs []int) int {
	// TODO: Run every job, join all workers, and sum every result.
	return 0
}

func main() {}
