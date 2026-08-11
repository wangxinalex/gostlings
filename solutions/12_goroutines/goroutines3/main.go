// Concept: passing values into goroutines
// Task: pass each label as an explicit goroutine argument

package main

import "sync"

func runWithArgs(labels []string) []string {
	results := make([]string, len(labels))
	var wg sync.WaitGroup
	for index := range labels {
		wg.Add(1)
		go func(index int, label string) {
			defer wg.Done()
			results[index] = label
		}(index, labels[index])
	}
	wg.Wait()
	return results
}

func main() {}
