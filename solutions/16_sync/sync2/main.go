// Concept: data races — concurrent writes to a shared map are unsafe
// Task: the program panics with a concurrent map write; add a mutex to make it safe
// Expected output: done
// Hint: map reads and writes must all be protected by the SAME mutex (Go doc: sync)

package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	m := make(map[int]int)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				mu.Lock()
				m[n*100+j] = j
				mu.Unlock()
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("done")
}
