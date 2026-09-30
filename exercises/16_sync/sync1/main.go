// Concept: sync.Mutex protects shared data from concurrent access
// Task: add a mutex to protect the counter so the total always equals 1000 (equal increments per goroutine)
// Expected output: total: 1000
// Hint: hold the lock across the whole read-modify-write, not just one step, and
//       release it on every path (Go doc: sync)
// Stuck?: lock the mutex before touching the counter and defer its Unlock.

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
