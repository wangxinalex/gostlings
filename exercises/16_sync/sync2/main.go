// Concept: data races — even a simple read-write to a shared map is unsafe
// Task: the program panics with a concurrent map write; add a mutex to make it safe
// Expected output: done
// Hint: map reads and writes must all be protected by the SAME mutex (Go doc: sync)

package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			// TODO: Protect this map write with a mutex to avoid the concurrent map write panic.
			m[n] = n * 10
		}(i)
	}

	wg.Wait()
	fmt.Println("done")
}
