// Concept: joining one batch before beginning the next
// Task: finish each batch of jobs before launching the following batch

package main

import (
	"fmt"
	"sync"
)

var runBatchJob = func(_ int, job int) string {
	return fmt.Sprintf("job %d done", job)
}

func runBatches(batches [][]int) [][]string {
	results := make([][]string, len(batches))
	for batchIndex, batch := range batches {
		batchResults := make([]string, len(batch))
		var wg sync.WaitGroup
		for jobIndex, job := range batch {
			wg.Add(1)
			go func(batchIndex, index, job int) {
				defer wg.Done()
				batchResults[index] = runBatchJob(batchIndex, job)
			}(batchIndex, jobIndex, job)
		}
		wg.Wait()
		results[batchIndex] = batchResults
	}
	return results
}

func main() {}
