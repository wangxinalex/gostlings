// Concept: sync.Mutex protects shared data from concurrent access
// Task: add a mutex to protect the counter so the total always equals 1000 (equal increments per goroutine)
// Expected output: total: 1000
// Hint: Lock before reading-modifying-writing, Unlock after; defer m.Unlock() is the common pattern (Go doc: sync)

package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var counter int
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// TODO: Protect the following line with the mutex.
			counter++
		}()
	}

	wg.Wait()
	fmt.Println("total:", counter)
}
