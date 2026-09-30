// Concept: passing values into goroutines
// Task: pass each label as an explicit goroutine argument
// Expected behavior: every label is returned, and empty input returns an empty slice.
// Hint: pass the loop values into the goroutine as parameters and read them from
//       the parameters inside the body instead of capturing the outer variables.
// Stuck?: a function literal can declare parameters and be called right away, so
//       each launch gets its own copy of the values.
// Note: explicit parameters work consistently before and after Go 1.22's
//       range-variable change.

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
