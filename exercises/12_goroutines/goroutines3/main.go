// Concept: passing values into goroutines
// Task: each goroutine should receive its own i as a parameter and print it
// Expected output: 0
// 1
// 2
// (any order)
// Hint: go func(n int) { fmt.Println(n) }(i) passes i explicitly. In Go 1.22+
//       loop variables are already per-iteration, but passing the value
//       documents intent and keeps the code correct on older Go versions
//       (Go Tour: Concurrency 1)

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		// TODO: Make the goroutine accept i as a parameter and print it.
		go func() {
			defer wg.Done()
		}()
	}

	wg.Wait()
}
