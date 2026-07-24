// Concept: sync.Once guarantees a function runs exactly once
// Task: the initConfig function should run only once even though it's called from multiple goroutines
// Expected output: config initialized
// running
// running
// Hint: var once sync.Once; once.Do(f) ensures f runs exactly once regardless of how many goroutines call it (Go doc: sync)

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
			once.Do(initConfig)
			fmt.Println("running")
		}()
	}

	wg.Wait()
}
