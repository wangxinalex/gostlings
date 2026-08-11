// Concept: sync.WaitGroup joins a known set of goroutines.
// Task: start count workers, wait for every worker, and return the completion count.
// Hint: call Add before each launch, defer Done in each goroutine, then Wait once.
package main

import "fmt"

var workerDone = func() {}

func waitForWorkers(count int) int {
	// TODO: Add and join every worker before returning its completion count.
	return 0
}

func main() { fmt.Println(waitForWorkers(3)) }
