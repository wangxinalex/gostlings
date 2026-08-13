// Concept: join each parallel batch before starting the next batch
// Task: run jobs inside each batch concurrently, but keep the batches sequential
// Expected behavior: preserve one result slice per batch; batch N+1 must not start
//                  until every job in batch N has finished
// Hint: use the outer loop for batch order and an inner loop for parallel jobs.
//       Create a result slice and a WaitGroup for each batch. Add before each go,
//       defer Done inside each worker, and Wait before the next outer-loop iteration.

package main

import (
	"fmt"
)

var runBatchJob = func(_ int, job int) string {
	return fmt.Sprintf("job %d done", job)
}

var onBatchStart = func(int) {}

func runBatches(batches [][]int) [][]string {
	// The pattern has two loops with two different concurrency rules:
	//   for each batch:                 // the outer loop is sequential
	//       start all jobs in this batch // jobs in one batch may run in parallel
	//       wait for this batch          // barrier / join before the next batch
	//       start the next batch
	//
	// Keep the two-level result shape: results[batchIndex][jobIndex].
	// Each worker should write only its own batchResults[jobIndex]; do not
	// concurrently append to one shared slice.
	// Suggested steps:
	//   1. results := make([][]string, len(batches))
	//   2. At the start of each batch, call onBatchStart(batchIndex)
	//   3. Create batchResults := make([]string, len(batch)) and a batch-local wg
	//   4. Start one goroutine per job, passing batchIndex, jobIndex, and job
	//   5. In each worker, defer wg.Done() and write batchResults[jobIndex]
	//   6. After wg.Wait(), store batchResults in results[batchIndex]
	// An empty batch needs no special case: wg.Wait() returns immediately.
	// TODO: Implement runBatches per the contract above.
	return nil
}

func main() {}
