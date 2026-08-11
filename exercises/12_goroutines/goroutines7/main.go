// Concept: joining one batch before beginning the next
// Task: finish each batch of jobs before launching the following batch
// Expected behavior: results keep their batch boundaries, and later batches start only after earlier batches finish.
// Hint: make a new WaitGroup for each batch, wait for it before the next outer-loop iteration, and keep one result slice per batch.

package main

import (
	"fmt"
)

var runBatchJob = func(_ int, job int) string {
	return fmt.Sprintf("job %d done", job)
}

func runBatches(batches [][]int) [][]string {
	// TODO: Run and join each batch in order, preserving its result boundary.
	return nil
}

func main() {}
