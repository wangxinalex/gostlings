// Concept: passing arguments to goroutines (the safe habit)
// Task: this program should print 0, 1, 2 but the goroutines print nothing — fix them to accept i as a parameter
// Expected output: 0
// 1
// 2
// (any order)
// Hint: go func(n int) { ... }(i) — passing values explicitly avoids surprises regardless of Go version (Go Tour: Concurrency 1)

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		// TODO: Make the goroutine accept i as a parameter so it prints the expected value.
		go func() {
			defer wg.Done()
		}()
	}

	wg.Wait()
}
