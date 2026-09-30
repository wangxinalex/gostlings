// Concept: sync.Once guarantees a function runs exactly once
// Task: the initConfig function should run only once even though it's called from multiple goroutines
// Expected output: config initialized
// running
// running
// running
// (any order)
// Hint: hand the initialization to a value that runs its function exactly once,
//       no matter how many goroutines call it (Go doc: sync)
// Stuck?: sync.Once and its Do method; callers still observe a single init.

package main

import (
	"fmt"
	"sync"
)

func initConfig() {
	fmt.Println("config initialized")
}

func main() {
	var once sync.Once
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// TODO: Use once.Do so initConfig runs exactly once.
			initConfig()
			fmt.Println("running")
		}()
	}

	wg.Wait()
}
