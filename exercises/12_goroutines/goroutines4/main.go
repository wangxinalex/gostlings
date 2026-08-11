// Concept: a basic sync.WaitGroup lifecycle
// Task: return one completion from every worker before runWorkers returns
// Expected behavior: count workers produce count completion strings, including zero workers.
// Hint: call Add before each launch, defer Done inside each worker, then call Wait before returning.

package main

func runWorkers(count int) []string {
	// TODO: Start count workers, join them, and return one completion per worker.
	return nil
}

func main() {}
