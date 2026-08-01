// Concept: data races — concurrent writes to a shared map are unsafe
// Task: the program panics with a concurrent map write; add a mutex to make it safe and verify all 10000 entries exist
// Expected output: entries: 10000
// Hint: map reads and writes must all be protected by the SAME mutex (Go doc: sync)

package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				// TODO: Protect this map write with a mutex to avoid the concurrent map write panic.
				m[n*100+j] = j
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("entries:", len(m))
}
