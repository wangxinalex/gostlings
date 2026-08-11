// Concept: sync.WaitGroup joins a known set of goroutines.
// Task: start count workers, wait for every worker, and return the completion count.
// Hint: call Add before each launch, defer Done in each goroutine, then Wait once.
package main

import (
	"fmt"
	"sync"
)

var workerDone = func() {}

func waitForWorkers(count int) int {
	if count <= 0 {
		return 0
	}
	var wg sync.WaitGroup
	completed := make(chan struct{}, count)
	wg.Add(count)
	for range count {
		go func() {
			defer wg.Done()
			workerDone()
			completed <- struct{}{}
		}()
	}
	wg.Wait()
	return len(completed)
}

func main() { fmt.Println(waitForWorkers(3)) }
