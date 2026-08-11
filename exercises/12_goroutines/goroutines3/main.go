// Concept: passing values into goroutines
// Task: pass each label as an explicit goroutine argument
// Expected behavior: every label is returned, and empty input returns an empty slice.
// Hint: use go func(index int, label string) { ... }(index, labels[index]) so the worker receives a value.

package main

import "sync"

func runWithArgs(labels []string) []string {
	results := make([]string, len(labels))
	var wg sync.WaitGroup
	for index := range labels {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			// TODO: Accept the label as an explicit goroutine argument and store it.
			results[index] = ""
		}(index)
	}
	wg.Wait()
	return results
}

func main() {}
