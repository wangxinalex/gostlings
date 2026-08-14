// Concept: closure capture in goroutines
// Task: make every goroutine keep the loop value it was created for
// Expected behavior: the returned slice contains every input label exactly once.
// Hint: create a new loop-local label before creating the closure; each closure needs its own value.
// Version note: label is deliberately declared outside the range loop and reused,
// so Go 1.22's per-iteration range variables do not fix this particular capture.

package main

import "sync"

func runLabels(labels []string) []string {
	results := make([]string, len(labels))
	tasks := make([]func() string, len(labels))
	var label string
	for index := range labels {
		label = labels[index]
		tasks[index] = func() string {
			// TODO: Capture a loop-local copy of label instead of this reused variable.
			return label
		}
	}

	var wg sync.WaitGroup
	for index, task := range tasks {
		wg.Add(1)
		go func(index int, task func() string) {
			defer wg.Done()
			results[index] = task()
		}(index, task)
	}
	wg.Wait()
	return results
}

func main() {}
