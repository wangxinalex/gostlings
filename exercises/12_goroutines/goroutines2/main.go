// Concept: synchronizing goroutines with sync.WaitGroup
// Task: add the missing WaitGroup calls so all three workers finish before main exits
// Expected output: worker 0 done
// worker 1 done
// worker 2 done
// (any order)
// Hint: call wg.Add before each go, wg.Done inside the goroutine, and wg.Wait at the end (Go Tour: Concurrency 1)

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		// TODO: Tell the WaitGroup about this goroutine.
		go func(n int) {
			// TODO: Mark this goroutine as done when it finishes.
			fmt.Println("worker", n, "done")
		}(i)
	}

	// TODO: Wait for all goroutines to finish.
}
