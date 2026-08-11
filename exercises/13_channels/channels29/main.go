// Concept: a pool can expose only the channel direction each caller needs.
// Task: return a send-only jobs handle and a receive-only results stream for a squaring pool.
// Expected behavior: callers close their jobs handle when done, then range results until the coordinator closes it.
// Hint: make jobs inside startPool, return it as chan<- int, and return results as <-chan int.
//       Workers range jobs; their exit acknowledgements let one coordinator close results exactly once.

package main

import "fmt"

var onPoolWorkerStart = func() {}
var onPoolWorkerExit = func() {}

func startPool(workers int) (chan<- int, <-chan int) {
	return nil, nil // TODO: expose directional handles and coordinate result closure
}

func main() {
	jobs, results := startPool(2)
	jobs <- 2
	jobs <- 3
	close(jobs)
	for result := range results {
		fmt.Println(result)
	}
}
