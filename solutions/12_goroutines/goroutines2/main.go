// Concept: closure capture in goroutines
// Task: make every goroutine keep the loop value it was created for

package main

import "sync"

func runLabels(labels []string) []string {
	results := make([]string, len(labels))
	tasks := make([]func() string, len(labels))
	for index := range labels {
		label := labels[index]
		tasks[index] = func() string {
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
